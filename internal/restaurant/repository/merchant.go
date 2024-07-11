package repository

var (
	ErrMerchantNotFound = errors.New("merchant not found")
)

type MerchantRepository interface {
	FindMerchantByID(ctx context.Context, id domain.MerchantID) (domain.Merchant, error)
	Update(ctx context.Context, merchant domain.Merchant) error
}