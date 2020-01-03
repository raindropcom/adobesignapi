package main

import (
	"bufio"
	"bytes"
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"

	"cloud.google.com/go/storage"
	"github.com/GoogleCloudPlatform/berglas/pkg/berglas"
	"github.com/dgrijalva/jwt-go"
	"github.com/jfcote87/esign"
	"github.com/jfcote87/esign/v2/envelopes"
	"github.com/jfcote87/esign/v2/model"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var pub *rsa.PublicKey

type metadata struct {
	userid string
	email  string
}

const rootPEM = `-----BEGIN CERTIFICATE-----
MIIC+TCCAeGgAwIBAgIJTNos05d3Em+zMA0GCSqGSIb3DQEBCwUAMBoxGDAWBgNV
BAMTD3JhaW5kLmF1dGgwLmNvbTAeFw0xOTA3MzAxNzIwNDFaFw0zMzA0MDcxNzIw
NDFaMBoxGDAWBgNVBAMTD3JhaW5kLmF1dGgwLmNvbTCCASIwDQYJKoZIhvcNAQEB
BQADggEPADCCAQoCggEBAPTW+Nfk8EoqeC0OdESca2Fra2rqynZ05/E/+ijQeAoi
iKs7FZmWeamZCR7Uj7R1jBNVfAghboNRoO/yRb6umctrrUdeiY2UZbVA4pRXO72G
t3kYu6zV/U60cCu9Aa1/NW75/gFj/yVKjIlQZmJw9J0O0iQ4kVLpQnP7UrjS3phM
HRok6C1+upnj3lD+EaMO5MlawmmPSehpb0/FR88ZqXH81ZLU7VjJlNIMGLohuNtL
zaMKeWeQBhoPeiLBpc+Fe/4SDYoLXygDb1zMQx9WLVAhYeA9XbscscKDU3miJM/6
8WUNnGFe+mWfcDc3ETcfJ7zZa20LJffJnSKZGw9mo1MCAwEAAaNCMEAwDwYDVR0T
AQH/BAUwAwEB/zAdBgNVHQ4EFgQUmDnLVmFHTeTc2qiB3BsjRmTaoNkwDgYDVR0P
AQH/BAQDAgKEMA0GCSqGSIb3DQEBCwUAA4IBAQCAsD3OIc7Fb1/zDTaDOVocY5/r
HL0/5WOHfnEcfBwg+gfSNJ3mQ3+LWvB0e278jNwqcdYQI8GotFDkHrD1FZ2DoplJ
wqLK1ONj+aviMgLXKh7xfOb14yLDg4tCIIcYSpJqQFiDu2mOa45bg3iHFPpswi+q
07evEP0SV0s0g7Y5bbUZ0JMiocKXZhFLCQ384bhPps9uT7KR+bqAGjAlh6QgR5gz
7idOEAwqT1hbnAJu1LKoV63IrLW1BEGMQxOj2lny2dWUKIN2jHJYJ0NdHy1IBNWt
kn45OyeFYW5/Urn1X4WGhkFHQXSxRTjRMeNHfEzGuCA8KWpR9oA/lDh6lepR
-----END CERTIFICATE-----`

var key []byte
var jwtKey = []byte("my_secret_key")

//verifyToken verifies jws signature
func verifyToken(token string) (bool, error) {
	block, _ := pem.Decode([]byte(rootPEM))
	if block == nil || block.Type != "CERTIFICATE" {
		log.Fatal("failed to decode PEM block containing public key")
	}
	var cert *x509.Certificate
	cert, _ = x509.ParseCertificate(block.Bytes)
	pub = cert.PublicKey.(*rsa.PublicKey)
	parts := strings.Split(token, ".")
	err := jwt.SigningMethodRS256.Verify(strings.Join(parts[0:2], "."), parts[2], pub)
	if err != nil {
		return false, nil
	}
	return true, nil
}

var sugar *zap.SugaredLogger

func initLogger() {
	logger, _ := zap.NewProduction()
	sugar = logger.Sugar()
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter() zapcore.WriteSyncer {
	var buffer bytes.Buffer
	writer := bufio.NewWriter(&buffer)
	return zapcore.AddSync(writer)
	//do something with buffered log
}

var awsAccessKeyID string
var awsSecretAccessKey string
var dsClientID string
var dsPrivateKey string

//initKey retrieves key from google kms using berglas
func initKey() {

	ctx := context.Background()
	berglasclient, err := berglas.New(ctx)
	if err != nil {
		sugar.Errorf("Error - %s", err)
	}
	bucket := "gke_cloud_run_keys"
	secret, err := berglasclient.Access(ctx, &berglas.AccessRequest{
		Bucket: "gke_cloud_run_keys",
		Object: "FILE_STORAGE_AES",
	})
	if err != nil {
		sugar.Errorf("Error - %s", err)
	}
	key = secret

	secret, err = berglasclient.Access(ctx, &berglas.AccessRequest{
		Bucket: bucket,
		Object: "DS_CLIENT_ID",
	})
	if err != nil {
		sugar.Errorf("Init - Failed to retrieve DS_CLIENT_ID key from bucket: "+bucket, ", Error - %s", err)
	}

	dsClientID = string(secret)

	secret, err = berglasclient.Access(ctx, &berglas.AccessRequest{
		Bucket: bucket,
		Object: "DS_PRIVATE_KEY",
	})
	if err != nil {
		sugar.Errorf("Init - Failed to retrieve DS_PRIVATE_KEY key from bucket: "+bucket, ", Error - %s", err)
	}

	dsPrivateKey = string(secret)
}

func main() {
	initLogger()
	defer sugar.Sync()
	initKey()
	http.HandleFunc("/send", send)
	http.HandleFunc("/consenturl", getConsentURL)
	http.HandleFunc("/status", getEnvelopeStatus)
	http.HandleFunc("/senderview", embeddSenderView)
	http.HandleFunc("/webhook/envelopestatus", handleWebHookEvelopeStatus)

	port := os.Getenv("PORT")
	log.Println(port)
	if port == "" {
		port = "443"
		log.Printf("Defaulting to port %s", port)
		sugar.Infof("Defaulting to port %s", port)
	}
	// Start HTTP server.
	sugar.Infof("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

type urls struct {
	SenderViewURL string `json:"senderviewurl,omnitempty"`
	ConsentURL    string `json:"consentURL,omnitempty"`
}

type envSummary struct {
	EnvelopeID     string `json:"envelopeid"`
	EnvelopeStatus string `json:"envelopestatus,omnitempty"`
	URLS           urls   `json:"urls,omnitempty"`
}

type envelopesStatuses struct {
	ImpersonatedUserGUID string                   `json:"impersonated_user_guid"`
	EnvelopeIdsRequest   model.EnvelopeIdsRequest `json:"envelopestatusrequest"`
}

type emailname struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type requestFile struct {
	Secret     string `json:"secret"`
	UUID       string `json:"uuid"`
	Generation int64  `json:"generation,omnitempty"`
	Name       string `json:"name"`
}

type requestBody struct {
	ImpersonatedUserGUID string        `json:"impersonated_user_guid"`
	Signers              []emailname   `json:"signers"`
	CCRecipients         []emailname   `json:"ccrecipients"`
	Files                []requestFile `json:"documents"`
}

func send(w http.ResponseWriter, r *http.Request) {
	if origin := r.Header.Get("Origin"); origin != "" {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Requested-With, FileKey")
		w.Header().Set("Access-Control-Max-Age", "3600")
	}
	// Stop here if its Preflighted OPTIONS request
	if r.Method == "OPTIONS" {
		sugar.Info("preflight options response")
		w.WriteHeader(http.StatusNoContent)
		return
	} else if r.Method != "POST" {
		sugar.Errorf("wrong http method")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// We can obtain the session token from the requests cookies, which come with every request
	c := r.Header.Get("Authorization")
	// log.Println("Secret key:", string(key))

	if c == "" {
		// If the cookie is not set, return an unauthorized status
		w.WriteHeader(http.StatusUnauthorized)
		// fmt.Println("no access token")
		sugar.Errorf("no access token")
		return
	}

	// Get the JWT string from the cookie
	tknStr := strings.Replace(c, `Bearer `, "", -1)
	isValid, err := verifyToken(tknStr)
	if err != nil {
		// log.Println(err)
		sugar.Errorf("cannot verify token")
	}

	if !isValid {
		w.WriteHeader(http.StatusBadRequest)
		// fmt.Println("invalid token")
		w.Write([]byte("invalid token"))
		sugar.Error("invalid token")
		return
	}

	sugar.Errorf("valid token")
	// Initialize a new instance of `Claims`

	claims := jwt.MapClaims{}

	// Parse the JWT string and store the result in `claims`.
	// Note that we are passing the key in this method as well. This method will return an error
	// if the token is invalid (if it has expired according to the expiry time we set on sign in),
	// or if the signature does not match
	_, err = jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return pub, nil
	})
	if err != nil {
		log.Println(err)
	}
	// ... error handling
	// do something with decoded claims
	// for key, val := range claims {
	// 	fmt.Printf("Key: %v, value: %v\n", key, val)
	// }

	num := claims["https://hasura.io/jwt/claims"]
	md, _ := num.(map[string]interface{})
	iid := md["x-hasura-instance-id"]
	uid := md["x-hasura-user-id"]
	email := md["x-hasura-email"]
	sugar.Infof("uid: ", uid, "email: ", email)
	// meta := metadata{uid.(string), email.(string)}

	sugar.Info("uid: ", uid, "email: ", email, "instance ID: ", iid)

	/*
		exp := claims["exp"]

			log.Println("unix exp time", exp)
			t := time.Now()
			log.Println("current time:", t.Unix())
			log.Println(t.Unix() == int64(exp.(float64)))
			//log.Println("token read")
	*/

	var projectID, bucket, name string

	bucket = "trevor100"
	projectID = "octopi"

	bucket = mustGetEnv("GOLANG_SAMPLES_BUCKET", bucket)
	projectID = mustGetEnv("GOLANG_SAMPLES_PROJECT_ID", projectID)
	ctx := context.Background()

	client, err := storage.NewClient(ctx)
	if err != nil {
		sugar.Errorf("cannot create client: ", err)
	}

	ed := &requestBody{}
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&ed)
	if err != nil {
		sugar.Error(err)
	}

	var documentArray []model.Document
	for i, req := range ed.Files {

		prmuuid, err := decrypt(key, req.Secret)
		prm := string(prmuuid[0]) //permission read write delete possibly resumable upload
		//log.Println("Permission:", prm)
		UUID := trimFirstRune(prmuuid)
		if err != nil {

			sugar.Errorf("invalid aes", err)
		}

		filename := "banana/" + iid.(string) + "/" + UUID

		if (string(prm) == "R") && (r.Method == http.MethodGet) { //read file with uuid
			sugar.Info("read file request")
			exist := objcheck(bucket, name)
			if exist == false {
				log.Println("cannot read object ")
				http.Error(w, "File not found.", 404)
				sugar.Errorf("object does not exist")
				return
			}
			var documentBase64 []byte

			_, buffer, err := readobj(client, bucket, filename, req.Generation)
			documentBase64 = buffer.Bytes()
			sugar.Infof("documentID: ", string(int(i+1)))
			doc := model.Document{
				DocumentBase64: documentBase64,
				Name:           req.Name,
				DocumentID:     string(int(i + 1)),
			}
			documentArray = append(documentArray, doc)
			if err != nil {
				sugar.Errorf("cannot read obj: ", err)
			}
		}
	}

	envsummary, err := handleSend(w, ed, documentArray, r, "sent")
	if err != nil {
		sugar.Error(err)
	}

	envInfo := envSummary{
		EnvelopeID:     envsummary.EnvelopeID,
		EnvelopeStatus: envsummary.Status,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(envInfo)
	return

}

func embeddSenderView(w http.ResponseWriter, r *http.Request) {
	if origin := r.Header.Get("Origin"); origin != "" {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Requested-With, FileKey")
		w.Header().Set("Access-Control-Max-Age", "3600")
	}
	// Stop here if its Preflighted OPTIONS request
	if r.Method == "OPTIONS" {
		sugar.Info("preflight options response")
		w.WriteHeader(http.StatusNoContent)
		return
	} else if r.Method != "POST" {
		sugar.Errorf("wrong http method")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// We can obtain the session token from the requests cookies, which come with every request
	c := r.Header.Get("Authorization")
	// log.Println("Secret key:", string(key))

	if c == "" {
		// If the cookie is not set, return an unauthorized status
		w.WriteHeader(http.StatusUnauthorized)
		// fmt.Println("no access token")
		sugar.Errorf("no access token")
		return
	}

	// Get the JWT string from the cookie
	tknStr := strings.Replace(c, `Bearer `, "", -1)
	isValid, err := verifyToken(tknStr)
	if err != nil {
		// log.Println(err)
		sugar.Errorf("cannot verify token")
	}

	if !isValid {
		w.WriteHeader(http.StatusBadRequest)
		// fmt.Println("invalid token")
		w.Write([]byte("invalid token"))
		sugar.Error("invalid token")
		return
	}

	sugar.Errorf("valid token")
	// Initialize a new instance of `Claims`

	claims := jwt.MapClaims{}

	// Parse the JWT string and store the result in `claims`.
	// Note that we are passing the key in this method as well. This method will return an error
	// if the token is invalid (if it has expired according to the expiry time we set on sign in),
	// or if the signature does not match
	_, err = jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return pub, nil
	})
	if err != nil {
		log.Println(err)
	}
	// ... error handling
	// do something with decoded claims
	// for key, val := range claims {
	// 	fmt.Printf("Key: %v, value: %v\n", key, val)
	// }

	num := claims["https://hasura.io/jwt/claims"]
	md, _ := num.(map[string]interface{})
	// iid := md["x-hasura-instance-id"]
	uid := md["x-hasura-user-id"]
	email := md["x-hasura-email"]
	sugar.Infof("uid: ", uid, "email: ", email)
	// meta := metadata{uid.(string), email.(string)}

	iid := md["x-hasura-instance-id"]

	sugar.Info("uid: ", uid, "email: ", email, "instance ID: ", iid)

	/*
		exp := claims["exp"]

			log.Println("unix exp time", exp)
			t := time.Now()
			log.Println("current time:", t.Unix())
			log.Println(t.Unix() == int64(exp.(float64)))
			//log.Println("token read")
	*/

	var projectID, bucket, name string

	bucket = "trevor100"
	projectID = "octopi"

	bucket = mustGetEnv("GOLANG_SAMPLES_BUCKET", bucket)
	projectID = mustGetEnv("GOLANG_SAMPLES_PROJECT_ID", projectID)
	ctx := context.Background()

	client, err := storage.NewClient(ctx)
	if err != nil {
		sugar.Errorf("cannot create client: ", err)
	}

	ed := &requestBody{}
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&ed)
	if err != nil {
		sugar.Error(err)
	}

	var documentArray []model.Document
	var docID int
	docID = 1
	for _, req := range ed.Files {
		log.Println(req.UUID)
		prmuuid, err := decrypt(key, req.Secret)
		prm := string(prmuuid[0])
		UUID := trimFirstRune(prmuuid)
		if err != nil {
			sugar.Errorf("invalid aes", err)
		}

		filename := "banana/" + iid.(string) + "/" + UUID

		if string(prm) == "R" { //read file with uuid
			exist := objcheck(bucket, name)
			if exist == false {

				http.Error(w, "File not found.", 404)
				sugar.Errorf("object does not exist")
				return
			}
			var documentBase64 []byte

			_, buffer, err := readobj(client, bucket, filename, req.Generation)
			documentBase64 = buffer.Bytes()
			doc := model.Document{
				DocumentBase64: documentBase64,
				Name:           req.Name,
				DocumentID:     strconv.Itoa(docID),
			}
			documentArray = append(documentArray, doc)
			if err != nil {
				sugar.Errorf("cannot read obj: ", err)
			}
		}

		docID = docID + 1
	}

	envsummary, err := handleSend(w, ed, documentArray, r, "created")
	if err != nil {
		sugar.Error(err)
	}
	viewURL, err := createSenderView(ed.ImpersonatedUserGUID, envsummary.EnvelopeID)
	if err != nil {
		sugar.Error(err)
	}
	viewInfo := urls{
		SenderViewURL: viewURL.URL,
	}

	envInfo := envSummary{
		EnvelopeID:     envsummary.EnvelopeID,
		EnvelopeStatus: envsummary.Status,
		URLS:           viewInfo,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(envInfo)
	return

}

func getConsentURL(w http.ResponseWriter, r *http.Request) {
	if origin := r.Header.Get("Origin"); origin != "" {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Requested-With, FileKey")
		w.Header().Set("Access-Control-Max-Age", "3600")
	}
	// Stop here if its Preflighted OPTIONS request
	if r.Method == "OPTIONS" {
		sugar.Info("preflight options response")
		w.WriteHeader(http.StatusNoContent)
		return
	} else if r.Method != "POST" {
		sugar.Errorf("wrong http method")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// We can obtain the session token from the requests cookies, which come with every request
	c := r.Header.Get("Authorization")
	// log.Println("Secret key:", string(key))

	if c == "" {
		// If the cookie is not set, return an unauthorized status
		w.WriteHeader(http.StatusUnauthorized)
		// fmt.Println("no access token")
		sugar.Errorf("no access token")
		return
	}

	// Get the JWT string from the cookie
	tknStr := strings.Replace(c, `Bearer `, "", -1)
	isValid, err := verifyToken(tknStr)
	if err != nil {
		// log.Println(err)
		sugar.Errorf("cannot verify token")
	}

	if !isValid {
		w.WriteHeader(http.StatusBadRequest)
		// fmt.Println("invalid token")
		w.Write([]byte("invalid token"))
		sugar.Error("invalid token")
		return
	}

	sugar.Errorf("valid token")
	// Initialize a new instance of `Claims`

	claims := jwt.MapClaims{}

	// Parse the JWT string and store the result in `claims`.
	// Note that we are passing the key in this method as well. This method will return an error
	// if the token is invalid (if it has expired according to the expiry time we set on sign in),
	// or if the signature does not match
	_, err = jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return pub, nil
	})
	if err != nil {
		sugar.Infof("jwt parse error: ", err)
	}
	// ... error handling
	// do something with decoded claims
	// for key, val := range claims {
	// 	fmt.Printf("Key: %v, value: %v\n", key, val)
	// }

	num := claims["https://hasura.io/jwt/claims"]
	md, _ := num.(map[string]interface{})
	// iid := md["x-hasura-instance-id"]
	uid := md["x-hasura-user-id"]
	email := md["x-hasura-email"]
	sugar.Infof("uid: ", uid, "email: ", email)

	ed := &requestBody{}
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&ed)
	if err != nil {
		sugar.Error(err)
	}
	consentStruct := urls{
		ConsentURL: getProxyURL(),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(consentStruct)
	return

}

func getEnvelopeStatus(w http.ResponseWriter, r *http.Request) {
	if origin := r.Header.Get("Origin"); origin != "" {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Requested-With, FileKey")
		w.Header().Set("Access-Control-Max-Age", "3600")
	}
	// Stop here if its Preflighted OPTIONS request
	if r.Method == "OPTIONS" {
		sugar.Info("preflight options response")
		w.WriteHeader(http.StatusNoContent)
		return
	} else if r.Method != "POST" {
		sugar.Errorf("wrong http method")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// We can obtain the session token from the requests cookies, which come with every request
	c := r.Header.Get("Authorization")
	// log.Println("Secret key:", string(key))

	if c == "" {
		// If the cookie is not set, return an unauthorized status
		w.WriteHeader(http.StatusUnauthorized)
		// fmt.Println("no access token")
		sugar.Errorf("no access token")
		return
	}

	// Get the JWT string from the cookie
	tknStr := strings.Replace(c, `Bearer `, "", -1)
	isValid, err := verifyToken(tknStr)
	if err != nil {
		// log.Println(err)
		sugar.Errorf("cannot verify token")
	}

	if !isValid {
		w.WriteHeader(http.StatusBadRequest)
		// fmt.Println("invalid token")
		w.Write([]byte("invalid token"))
		sugar.Error("invalid token")
		return
	}

	sugar.Errorf("valid token")
	// Initialize a new instance of `Claims`

	claims := jwt.MapClaims{}

	// Parse the JWT string and store the result in `claims`.
	// Note that we are passing the key in this method as well. This method will return an error
	// if the token is invalid (if it has expired according to the expiry time we set on sign in),
	// or if the signature does not match
	_, err = jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return pub, nil
	})
	if err != nil {
		sugar.Infof("jwt parse error: ", err)
	}
	// ... error handling
	// do something with decoded claims
	// for key, val := range claims {
	// 	fmt.Printf("Key: %v, value: %v\n", key, val)
	// }

	num := claims["https://hasura.io/jwt/claims"]
	md, _ := num.(map[string]interface{})
	// iid := md["x-hasura-instance-id"]
	uid := md["x-hasura-user-id"]
	email := md["x-hasura-email"]
	sugar.Infof("uid: ", uid, "email: ", email)
	envs := &envelopesStatuses{}
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&envs)
	if err != nil {
		sugar.Error(err)
	}
	ctx := context.TODO()
	credential, err := getCredentials(envs.ImpersonatedUserGUID)
	if err != nil {
		sugar.Error(err)
		http.Error(w, "Unauthorized Credentials, Consent Required", http.StatusUnauthorized)
	}
	sv := envelopes.New(credential)
	envsInfo, err := sv.ListStatus(&envs.EnvelopeIdsRequest).Do(ctx)
	if err != nil {
		sugar.Error(err)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(envsInfo)
	return
}

const charSet = "UTF-8"

var secret string

func handleSend(w http.ResponseWriter, requestBody *requestBody, documents []model.Document, r *http.Request, status string) (*model.EnvelopeSummary, error) {
	credential, err := getCredentials(requestBody.ImpersonatedUserGUID)
	if err != nil {
		sugar.Error(err)
		http.Error(w, "Unauthorized Credentials, Consent Required", http.StatusUnauthorized)
		return nil, err
	}
	sv := envelopes.New(credential)
	var signers []model.Signer
	for i, signer := range requestBody.Signers {

		u1 := uuid.NewV4()
		signerBody := model.Signer{

			EmailNotification: nil,

			RecipientID: u1.String(),
		}
		signers = append(signers, signerBody)
		signers[i].Email = signer.Email
		signers[i].Name = signer.Name
	}
	env := &model.EnvelopeDefinition{
		EmailSubject: "[Go eSignagure SDK] - Please sign this doc",
		EmailBlurb:   "Please sign this test document",
		Status:       status,
		Documents:    documents,
		Recipients: &model.Recipients{
			Signers: signers,
		},
	}
	envSummary, err := sv.Create(env).Do(context.Background())
	return envSummary, err
}

func getCredentials(userProxy string) (*esign.OAuth2Credential, error) {

	privateKey2 := "-----BEGIN RSA PRIVATE KEY-----" + "\n" + dsPrivateKey + "\n" + "-----END RSA PRIVATE KEY-----"

	cfg := &esign.JWTConfig{
		IntegratorKey: dsClientID,
		KeyPairID:     "b6756234-b625-4fe1-b5ea-fd336a3c9b32",
		PrivateKey:    privateKey2,
		IsDemo:        true,
	}

	apiUserName := userProxy
	//this is the impersonated user guid
	credential, err := cfg.Credential(apiUserName, nil, nil)
	if err != nil {
		log.Fatalf("create credential error: %v", err)
	}

	return credential, err
}

func getcfg(proxyID string) *esign.JWTConfig {

	privateKey := "-----BEGIN RSA PRIVATE KEY-----" + "\n" + dsPrivateKey + "\n" + "-----BEGIN RSA PRIVATE KEY-----"
	cfg := &esign.JWTConfig{
		IntegratorKey: dsClientID,
		AccountID:     "0bf5720a-2158-4352-abac-658daf277e35",
		KeyPairID:     "b6756234-b625-4fe1-b5ea-fd336a3c9b32",
		PrivateKey:    privateKey,
		IsDemo:        true,
	}
	return cfg
}

func getProxyURL() string {

	cfg := getcfg("")

	consentURL := cfg.UserConsentURL("http://localhost:3000/docusign/callback")
	// Redirect user to consent page.
	return consentURL
}

// func getProxyCredentials(code string) *esign.OAuth2Credential {
// 	ctx := context.TODO()
// 	cfg := getcfg("")
// 	// Enter code returned to redirectURL.

// 	if _, err := fmt.Scan(&code); err != nil {
// 		log.Fatal(err)
// 	}
// 	credential, err := cfg.Exchange(ctx, code)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fl, err := folders.New(credential).List().Do(ctx)
// 	if err != nil {
// 		log.Fatalf("Folder list error: %v", err)
// 	}
// 	for _, fld := range fl.Folders {
// 		fmt.Printf("%s: %s", fld.Name, fld.FolderID)
// 	}
// 	return credential
// }

func createSenderView(proxyGUID string, envelopeID string) (*model.ViewURL, error) {
	credential, err := getCredentials(proxyGUID)
	if err != nil {
		sugar.Error(err)
		return nil, err
	}
	returnURLReq := model.ReturnURLRequest{
		ReturnURL: "http://localhost:3000/docusign/callback",
	}
	ctx := context.TODO()
	viewURL, err := envelopes.New(credential).ViewsCreateSender(envelopeID, &returnURLReq).Do(ctx)
	return viewURL, err
}

func readobj(client *storage.Client, bucket, object string, gen int64) (*storage.ObjectAttrs, bytes.Buffer, error) {
	ctx := context.Background()
	var b bytes.Buffer
	// [START download_file]
	var err error
	var rc *storage.Reader
	if gen == 0 {
		rc, err = client.Bucket(bucket).Object(object).NewReader(ctx)

	} else {
		rc, err = client.Bucket(bucket).Object(object).Generation(gen).NewReader(ctx)
	}
	if err != nil {
		sugar.Error(err)
		return nil, b, err
	}

	bh := client.Bucket(bucket)
	obj := bh.Object(object)
	objAttrs, err := obj.Attrs(ctx)

	//w.Header().Set("Content-Length", strconv.FormatInt(objAttrs.Size, 10))

	bwriter := bufio.NewWriter(&b)
	a, err := io.Copy(bwriter, rc)
	defer rc.Close()
	log.Println(strconv.FormatInt(a, 10))
	return objAttrs, b, err
}

func trimFirstRune(s string) string {
	_, i := utf8.DecodeRuneInString(s)
	return s[i:]
}

func decrypt(key []byte, text string) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	decodedMsg, err := base64.StdEncoding.DecodeString(addBase64Padding(text))
	if err != nil {
		return "", err
	}

	if (len(decodedMsg) % aes.BlockSize) != 0 {
		return "", errors.New("blocksize must be multipe of decoded message length")
	}

	iv := decodedMsg[:aes.BlockSize]
	msg := decodedMsg[aes.BlockSize:]

	cfb := cipher.NewCBCDecrypter(block, iv)
	cfb.CryptBlocks(msg, msg)
	unpadMsg, err := Unpad(msg)
	if err != nil {
		return "", err
	}

	return string(unpadMsg), nil
}

//Unpad unpad aes message
func Unpad(src []byte) ([]byte, error) {

	length := len(src)
	unpadding := int(src[length-1])

	if unpadding > length {
		return nil, errors.New("unpad error. This could happen when incorrect encryption key is used")
	}

	return src[:(length - unpadding)], nil
}

func addBase64Padding(value string) string {
	m := len(value) % 4
	if m != 0 {
		value += strings.Repeat("=", 4-m)
	}

	return value
}

func objcheck(bucket, name string) bool {
	client, err := google.DefaultClient(oauth2.NoContext,
		"https://www.googleapis.com/auth/devstorage.full_control")
	if err != nil {
		//log.Fatal(err)
		sugar.Error(err)
	}
	req, err := http.NewRequest("GET", "https://www.googleapis.com/storage/v1/b/"+bucket+"/"+"o/"+url.QueryEscape(name), nil)
	//if the object doesnt exist get returns a 404 status error
	resp, err := client.Do(req)
	var exist bool
	exist = true
	log.Println(resp.StatusCode)
	if resp.StatusCode == 404 {
		exist = false
	}

	if err != nil {
		exist = false
	}
	return exist
}

func mustGetEnv(envKey, defaultValue string) string {
	val := os.Getenv(envKey)
	if val == "" {
		val = defaultValue
	}
	if val == "" {
		log.Fatalf("%q should be set", envKey)
	}
	return val
}

func handleWebHookEvelopeStatus(w http.ResponseWriter, r *http.Request) {
	if origin := r.Header.Get("Origin"); origin != "" {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Requested-With, FileKey")
		w.Header().Set("Access-Control-Max-Age", "3600")
	}
	// Stop here if its Preflighted OPTIONS request
	if r.Method == "OPTIONS" {
		sugar.Info("preflight options response")
		w.WriteHeader(http.StatusNoContent)
		return
	} else if r.Method != "POST" {
		sugar.Errorf("wrong http method")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// We can obtain the session token from the requests cookies, which come with every request
	c := r.Header.Get("Authorization")
	// log.Println("Secret key:", string(key))

	if c == "" {
		// If the cookie is not set, return an unauthorized status
		w.WriteHeader(http.StatusUnauthorized)
		// fmt.Println("no access token")
		sugar.Errorf("no access token")
		return
	}

	// Get the JWT string from the cookie
	tknStr := strings.Replace(c, `Bearer `, "", -1)
	isValid, err := verifyToken(tknStr)
	if err != nil {
		// log.Println(err)
		sugar.Errorf("cannot verify token")
	}

	if !isValid {
		w.WriteHeader(http.StatusBadRequest)
		// fmt.Println("invalid token")
		w.Write([]byte("invalid token"))
		sugar.Error("invalid token")
		return
	}

	sugar.Errorf("valid token")
	// Initialize a new instance of `Claims`

	claims := jwt.MapClaims{}

	// Parse the JWT string and store the result in `claims`.
	// Note that we are passing the key in this method as well. This method will return an error
	// if the token is invalid (if it has expired according to the expiry time we set on sign in),
	// or if the signature does not match
	_, err = jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return pub, nil
	})
	if err != nil {
		sugar.Infof("jwt parse error: ", err)
	}
	// ... error handling
	// do something with decoded claims
	// for key, val := range claims {
	// 	fmt.Printf("Key: %v, value: %v\n", key, val)
	// }

	num := claims["https://hasura.io/jwt/claims"]
	md, _ := num.(map[string]interface{})
	// iid := md["x-hasura-instance-id"]
	uid := md["x-hasura-user-id"]
	email := md["x-hasura-email"]
	sugar.Infof("uid: ", uid, "email: ", email)

	ed := &requestBody{}
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&ed)
	if err != nil {
		sugar.Error(err)
	}
	consentStruct := urls{
		ConsentURL: getProxyURL(),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(consentStruct)
	return
}

func makeAgreement() {
	ctx := context.TODO()
	service := &DefaultApiService{}
	authorization := "string"
	agreementInfo := AgreementInfo{}
	localVarOptionals := &CreateAgreementOpts{}
	service.CreateAgreement(ctx, authorization, agreementInfo, localVarOptionals)
}
