package repository

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/fazilnbr/project-workey/pkg/domain"
)

func TestInsertUser(t *testing.T) {
	arg := domain.Login{
		UserName: "sethu",
		Password: "12345",
	}

	id, err := testuserrepo.InsertUser(arg)
	require.NoError(t, err)
	require.NotEmpty(t, id)
	require.NotZero(t, id)

}
