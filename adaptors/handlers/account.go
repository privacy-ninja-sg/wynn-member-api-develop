package handlers

import (
	"encoding/hex"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"wynn-member-api/ent"
	"wynn-member-api/ent/user"
	"wynn-member-api/internal/core/handlers"
	"wynn-member-api/internal/core/repositories"
	"wynn-member-api/internal/core/services"
	"wynn-member-api/pkg/errs"
	"wynn-member-api/pkg/utils"
)

type accountHandler struct {
	db          *ent.Client
	authService services.AuthService
	userRepo    repositories.UserRepository
}

type changePwdReq struct {
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

func NewAccountHandler(entCli *ent.Client, authServ services.AuthService, userRepo repositories.UserRepository) handlers.AccountHandler {
	return &accountHandler{
		db:          entCli,
		authService: authServ,
		userRepo:    userRepo,
	}
}

func (h accountHandler) Info(c *fiber.Ctx) error {
	tokenClaims := c.Locals("tokenPayload").(*utils.TokenMetadata)

	// get account info
	userData, err := h.db.User.Query().
		Where(user.ID(tokenClaims.Uid)).
		WithBanks().
		WithChannel().
		First(c.Context())
	if err != nil {
		if ent.IsNotFound(err) {
			return errs.NewNotFoundError(errs.NOT_FOUND)
		}
	}
	return utils.NewResponse(c, "ok", http.StatusOK, userData)
}

func (h accountHandler) ChangePassword(c *fiber.Ctx) error {
	tokenClaims := c.Locals("tokenPayload").(*utils.TokenMetadata)
	// validate zone
	var reqBody changePwdReq
	err := utils.ValidateBody(c, &reqBody)
	if err != nil {
		return err
	}

	if reqBody.NewPassword == reqBody.OldPassword {
		return utils.NewResponse(c, "ok", http.StatusOK, "successfully")
	}

	// validate old password
	err = h.authService.ValidatePassword(c.Context(), reqBody.OldPassword, tokenClaims.Uid)
	if err != nil {
		return errs.NewBadRequestError("please recheck your old password again")
	}

	// encrypt pwd
	newPwd, err := h.authService.EncryptText(reqBody.NewPassword)
	if err != nil {
		return errs.NewBadRequestError("please verify your new password again")
	}

	cipherPwd := hex.EncodeToString(newPwd)

	_, err = h.userRepo.UpdatePassword(c.Context(), tokenClaims.Uid, cipherPwd)
	if err != nil {
		return errs.NewInternalServerError()
	}

	return utils.NewResponse(c, "ok", http.StatusOK, "successfully")
}
