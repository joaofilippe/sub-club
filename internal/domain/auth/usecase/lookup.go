package usecase

import (
	"context"
	"fmt"
	"strings"
	"sync"

	"github.com/joaofilippe/subclub/internal/domain/account"
	"github.com/joaofilippe/subclub/internal/domain/accountdomain"
	authdomain "github.com/joaofilippe/subclub/internal/domain/auth"
	authmodel "github.com/joaofilippe/subclub/internal/domain/auth/model"
)

type LookupUseCase struct {
	accountRepo       account.Repository
	accountDomainRepo accountdomain.Repository
	tenantRepo        authdomain.TenantRepository
}

func NewLookupUseCase(
	accountRepo account.Repository,
	accountDomainRepo accountdomain.Repository,
	tenantRepo authdomain.TenantRepository,
) *LookupUseCase {
	return &LookupUseCase{
		accountRepo:       accountRepo,
		accountDomainRepo: accountDomainRepo,
		tenantRepo:        tenantRepo,
	}
}

func (uc *LookupUseCase) Execute(ctx context.Context, input authmodel.LookupInput) ([]authmodel.AccountInfo, error) {
	parts := strings.SplitN(input.Email, "@", 2)
	if len(parts) == 2 {
		info, err := uc.accountDomainRepo.GetByDomain(ctx, parts[1])
		if err == nil {
			return []authmodel.AccountInfo{{Slug: info.Slug, Name: info.Name}}, nil
		}
	}

	accounts, err := uc.accountRepo.List(ctx)
	if err != nil {
		return nil, fmt.Errorf("listing accounts: %w", err)
	}

	type hit struct {
		info  authmodel.AccountInfo
		found bool
	}

	results := make(chan hit, len(accounts))
	var wg sync.WaitGroup

	for _, acc := range accounts {
		wg.Add(1)
		go func(slug, name string) {
			defer wg.Done()
			exists, err := uc.tenantRepo.UserExistsByEmailAndSlug(ctx, slug, input.Email)
			if err != nil || !exists {
				results <- hit{}
				return
			}
			results <- hit{info: authmodel.AccountInfo{Slug: slug, Name: name}, found: true}
		}(acc.Slug, acc.Name)
	}

	wg.Wait()
	close(results)

	var matches []authmodel.AccountInfo
	for r := range results {
		if r.found {
			matches = append(matches, r.info)
		}
	}
	return matches, nil
}
