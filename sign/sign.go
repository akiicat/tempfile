package sign

import (
        "fmt"
        "time"
        "encoding/base64"
        "golang.org/x/oauth2/google"
        "cloud.google.com/go/storage"
      )

func Sign(serviceAccount, bucketName, objectName, method string, expireTime int) (string, error) {
  jsonKey, err := base64.StdEncoding.DecodeString(serviceAccount)
	if err != nil {
		return "", fmt.Errorf("service account decode error: %v", err)
	}

  conf, err := google.JWTConfigFromJSON(jsonKey)
  if err != nil {
    return "", fmt.Errorf("service key is not valid, google.JWTConfigFromJSON: %v", err)
  }

  // https://github.com/googleapis/google-cloud-go/blob/master/storage/storage.go#L157
  opts := &storage.SignedURLOptions{
    Scheme:         storage.SigningSchemeV4,
    Method:         method,
    GoogleAccessID: conf.Email,
    PrivateKey:     conf.PrivateKey,
    Expires:        time.Now().Add(time.Duration(expireTime) * time.Minute),
  }

  url, err := storage.SignedURL(bucketName, objectName, opts)
  if err != nil {
    return "", fmt.Errorf("Unable to generate a signed URL: %v", err)
  }

  return url, nil
}



