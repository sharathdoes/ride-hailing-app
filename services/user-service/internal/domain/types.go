package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)
type Role string

const (
	RoleCaptain Role = "captain"
	RoleUser Role = "user"
)

type User struct {
    ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
    Name         string             `bson:"name"          json:"name"`
    Phone        string             `bson:"phone"         json:"phone"`
    Email        string             `bson:"email"         json:"email"`
    Gender       string             `bson:"gender"        json:"gender"`
    Role         Role               `bson:"role"          json:"role"`
    ProfileImage string             `bson:"profile_image" json:"profile_image"`
    IsActive     bool               `bson:"is_active"     json:"is_active"`
    IsBlocked    bool               `bson:"is_blocked"    json:"is_blocked"`
    CreatedAt    time.Time          `bson:"created_at"    json:"created_at"`
    UpdatedAt    time.Time          `bson:"updated_at"    json:"updated_at"`
}

type Driver struct {
    ID             primitive.ObjectID `bson:"_id,omitempty"    json:"id"`
    UserID         primitive.ObjectID `bson:"user_id"          json:"user_id"`
    Vehicle        Vehicle            `bson:"vehicle"          json:"vehicle"`
    Documents 	   Documents     	  `bson:"documents"  json:"documents"`
    IsAvailable    bool               `bson:"is_available"     json:"is_available"`
    IsVerified     bool               `bson:"is_verified"      json:"is_verified"`
    Rating         float64            `bson:"rating"           json:"rating"`
    TotalRides     int                `bson:"total_rides"      json:"total_rides"`
}

type Vehicle struct {
    Type   string `bson:"type"   json:"type"`    // "bike", "auto", "car"
	Model string `bson:"model" json:"model"` // "Honda Activa", "Bajaj Auto"
    Number string `bson:"number" json:"number"`
}

type Documents struct {
    CollegeID   string `bson:"college_id"   json:"college_id"`   // student ID / roll number
    Department  string `bson:"department"   json:"department"`
    Year        int    `bson:"year"         json:"year"`
    IDCardImage string `bson:"id_card_image" json:"id_card_image"` // photo of college ID
    VerifiedAt  *time.Time `bson:"verified_at" json:"verified_at"` // nil until admin verifies
}