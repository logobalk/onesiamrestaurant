package usecase

type (
	ChangeMerhcantNameInput struct {
		MerchantID string
		Name string
	}

	ChangeMerhcantNameOutput struct {}
)

var (
	ErrMerchantNotActive = errors.New("merchant not active")
)

type MerchantUserUseCase interface {
	ChangeMerhcantName(ctx context.Context, input ChangeMerhcantNameInput) (ChangeMerhcantNameOutput, error)
}

type MerchantUserUseCaseImpl struct {
	merchantService service.MerchantService
	merchantRepository repository.MerchantRepository
}


func (u *MerchantUserUseCaseImpl) ChangeMerhcantName(
	ctx context.Context,
	input ChangeMerhcantNameInput,
) (ChangeMerhcantNameOutput, error) {
	mid, err := domain.NewMerchantID(input.MerchantID)
	if err != nil {
		var zero ChangeMerhcantNameOutput
		return zero, err
	}

	active, err := u.merchantService.IsMerchantActive(ctx, mid)
	if err != nil {
		var zero ChangeMerhcantNameOutput
		return zero, err
	}

	if !active {
		var zero ChangeMerhcantNameOutput
		return zero, ErrMerchantNotActive
	}

	merchant, err := u.merchantRepository.FindMerchantByID(ctx, mid)
	if err != nil {
		var zero ChangeMerhcantNameOutput
		return zero, err
	}

	err := merchant.ChangeName(input.Name)
	if err != nil {
		var zero ChangeMerhcantNameOutput
		return zero, err
	}

	err = u.merchantRepository.Update(ctx, merchant)
	if err != nil {
		var zero ChangeMerhcantNameOutput
		return zero, err
	}

	return ChangeMerhcantNameOutput{}, nil
}