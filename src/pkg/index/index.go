package index

import (
	"context"

	api "github.com/ibad69/go-devxcoaches/src/pkg"
)

type Repo interface {
	AddStudent(ctx context.Context, student api.Student) (api.Student, error)
	// GetStudent(ctx context.Context) (api.Student, error)
	GetClass(ctx context.Context) ([]api.Class, error)
	GetStudents(ctx context.Context) ([]api.Student, error)
	AddClass(ctx context.Context, class api.Class) (api.Class, error)
	AddSubject(ctx context.Context, subject api.Subject) (api.Subject, error)
	GetSubjects(ctx context.Context) ([]api.Subject, error)
	GetBatches(ctx context.Context) ([]api.Batch, error)
	// GetPeriods(ctx context.Context) (api.Period, error)
	AddPeriods(ctx context.Context, period api.Period) (api.Period, error)
	AddYear(ctx context.Context, year api.Year) (api.Year, error)
	AddBatch(ctx context.Context, Batch api.Batch) (api.Batch, error)
	GetYearWithMonths(ctx context.Context, yearId string) ([]api.YearMonthJoin, error)
	GetYears(ctx context.Context) ([]api.Year, error)
}

type Service interface {
	AddStudent(ctx context.Context, student api.Student) (api.Student, error)
	// GetStudent(ctx context.Context) (api.Student, error)
	GetClass(ctx context.Context) ([]api.Class, error)
	AddClass(ctx context.Context, class api.Class) (api.Class, error)
	AddSubject(ctx context.Context, subject api.Subject) (api.Subject, error)
	GetSubjects(ctx context.Context) ([]api.Subject, error)
	GetBatches(ctx context.Context) ([]api.Batch, error)
	// GetPeriods(ctx context.Context) (api.Period, error)
	AddPeriods(ctx context.Context, period api.Period) (api.Period, error)
	AddYear(ctx context.Context, year api.Year) (api.Year, error)
	GetStudents(ctx context.Context) ([]api.Student, error)
	AddBatch(ctx context.Context, Batch api.Batch) (api.Batch, error)
	GetYearWithMonths(ctx context.Context, yearId string) ([]api.YearMonthJoin, error)
	GetYears(ctx context.Context) ([]api.Year, error)
}

type repoService struct {
	repo Repo
}

func New(r Repo) Service {
	return &repoService{repo: r}
}

func (r *repoService) AddStudent(ctx context.Context, student api.Student) (api.Student, error) {
	return r.repo.AddStudent(ctx, student)
}

func (r *repoService) AddClass(ctx context.Context, Class api.Class) (api.Class, error) {
	return r.repo.AddClass(ctx, Class)
}

func (r *repoService) AddPeriods(ctx context.Context, Periods api.Period) (api.Period, error) {
	return r.repo.AddPeriods(ctx, Periods)
}

func (r *repoService) AddYear(ctx context.Context, Year api.Year) (api.Year, error) {
	return r.repo.AddYear(ctx, Year)
}

func (r *repoService) AddSubject(ctx context.Context, Subject api.Subject) (api.Subject, error) {
	return r.repo.AddSubject(ctx, Subject)
}

func (r *repoService) AddBatch(ctx context.Context, Batch api.Batch) (api.Batch, error) {
	return r.repo.AddBatch(ctx, Batch)
}

func (r *repoService) GetStudents(ctx context.Context) ([]api.Student, error) {
	return r.repo.GetStudents(ctx)
}

func (r *repoService) GetClass(ctx context.Context) ([]api.Class, error) {
	return r.repo.GetClass(ctx)
}

func (r *repoService) GetBatches(ctx context.Context) ([]api.Batch, error) {
	return r.repo.GetBatches(ctx)
}

func (r *repoService) GetSubjects(ctx context.Context) ([]api.Subject, error) {
	return r.repo.GetSubjects(ctx)
}

func (r *repoService) GetYearWithMonths(
	ctx context.Context,
	yearId string,
) ([]api.YearMonthJoin, error) {
	return r.repo.GetYearWithMonths(ctx, yearId)
}

func (r *repoService) GetYears(ctx context.Context) ([]api.Year, error) {
	return r.repo.GetYears(ctx)
}
