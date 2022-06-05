package main

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/gin-gonic/gin"
)

func validatePostMhsHandler(c *gin.Context) (*Mahasiswa, error) {
	npm := c.PostForm("npm")
	if npm == "" {
		err := errors.New("npm cannot be empty")
		return nil, err
	}

	nama := c.PostForm("nama")
	if nama == "" {
		err := errors.New("nama cannot be empty")
		return nil, err
	}

	newMhs := Mahasiswa{
		NPM:  npm,
		Nama: nama,
	}

	return &newMhs, nil
}

func createMhs(newMhs Mahasiswa) (*Mahasiswa, error) {
	mhs, err := json.Marshal(newMhs)
	if err != nil {
		return nil, err
	}

	err = mhsDB.Set(newMhs.NPM, mhs, time.Duration(0)).Err()
	if err != nil {
		return nil, err
	}

	return &newMhs, nil
}
