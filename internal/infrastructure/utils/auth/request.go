package auth

import (
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"strconv"

	"github.com/ahmetilboga2004/internal/application/middlewares"
)

func GetUserIDFromContext(r *http.Request) (uint, error) {
	rawUserID := r.Context().Value(middlewares.UserIDKey)
	fmt.Println("rawUserId: ", rawUserID)
	fmt.Println(reflect.TypeOf(rawUserID))

	if rawUserID == nil {
		return 0, errors.New("user ID not found in context")
	}

	if userIDStr, ok := rawUserID.(string); ok {
		userID, err := strconv.ParseUint(userIDStr, 10, 64)
		if err != nil {
			return 0, err
		}
		return uint(userID), nil
	}

	if userID, ok := rawUserID.(uint); ok {
		return userID, nil
	}

	if userID, ok := rawUserID.(uint64); ok {
		return uint(userID), nil
	}

	if userID, ok := rawUserID.(float64); ok {
		return uint(userID), nil
	}

	return 0, errors.New("unsupported user ID type in context")
}
