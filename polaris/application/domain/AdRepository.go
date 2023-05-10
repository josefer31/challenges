package domain

type AdRepository interface {
	Save(ad Ad) Ad
}
