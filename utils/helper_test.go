package utils

import (
	"reflect"
	"testing"
	"time"

	"github.com/golang-jwt/jwt"
)

func TestValidateCookie(t *testing.T) {

	expectedClaims := &jwt.StandardClaims{
		Audience:  "",
		ExpiresAt: time.Now().Add(time.Hour).Unix(),
	}

	claims, err := ValidateCookie("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDkyMDUwMTIsImlzcyI6IjEifQ.mOX5liW40Rs30qwczR_tAqYw-dffjOJ-3SndrgOm23s")

	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}

	if !reflect.DeepEqual(claims, expectedClaims) {
		t.Errorf("Expected %v, got %v", expectedClaims, claims)
	}
}
