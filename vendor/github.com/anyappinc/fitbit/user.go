package fitbit

import (
	"context"
	"encoding/json"
	"net/url"
	"time"
)

type (
	rawBadge struct {
		EncodedID               string        `json:"encodedId"`
		Name                    string        `json:"name"`
		ShortName               string        `json:"shortName"`
		Description             string        `json:"description"`
		ShortDescription        string        `json:"shortDescription"`
		MobileDescription       string        `json:"mobileDescription"`
		MarketingDescription    string        `json:"marketingDescription"`
		EarnedMessage           string        `json:"earnedMessage"`
		Value                   int64         `json:"value"`
		Unit                    string        `json:"unit"`
		DateTime                string        `json:"dateTime"`
		TimesAchieved           int64         `json:"timesAchieved"`
		Cheers                  []interface{} `json:"cheers"`
		Category                string        `json:"category"`
		BadgeType               string        `json:"badgeType"`
		BadgeGradientEndColor   string        `json:"badgeGradientEndColor"`
		BadgeGradientStartColor string        `json:"badgeGradientStartColor"`
		Image50Px               string        `json:"image50px"`
		Image75Px               string        `json:"image75px"`
		Image100Px              string        `json:"image100px"`
		Image125Px              string        `json:"image125px"`
		Image300Px              string        `json:"image300px"`
		ShareImage640Px         string        `json:"shareImage640px"`
		ShareText               string        `json:"shareText"`
	}

	// Badge represents a user's badge.
	Badge struct {
		EncodedID               string
		Name                    string
		ShortName               string
		Description             string
		ShortDescription        string
		MobileDescription       string
		MarketingDescription    string
		EarnedMessage           string
		Value                   int64
		Unit                    string
		DateTime                *time.Time
		TimesAchieved           int64
		Cheers                  []interface{}
		Category                string
		BadgeType               string
		BadgeGradientEndColor   string
		BadgeGradientStartColor string
		Image50Px               *url.URL
		Image75Px               *url.URL
		Image100Px              *url.URL
		Image125Px              *url.URL
		Image300Px              *url.URL
		ShareImage640Px         *url.URL
		ShareText               string
	}

	// Features represents user's features.
	Features struct {
		ExerciseGoal bool `json:"exerciseGoal"`
	}

	rawProfile struct {
		User struct {
			EncodedID                string    `json:"encodedId"`
			DisplayName              string    `json:"displayName"`
			DisplayNameSetting       string    `json:"displayNameSetting"`
			FullName                 string    `json:"fullName"`
			FirstName                string    `json:"firstName"`
			LastName                 string    `json:"lastName"`
			Gender                   string    `json:"gender"`
			DateOfBirth              string    `json:"dateOfBirth"`
			Age                      int64     `json:"age"`
			PhoneNumber              string    `json:"phoneNumber"`
			Country                  string    `json:"country"`
			State                    string    `json:"state"`
			City                     string    `json:"city"`
			Timezone                 string    `json:"timezone"`
			OffsetFromUTCMillis      int64     `json:"offsetFromUTCMillis"`
			Height                   float64   `json:"height"`
			Weight                   float64   `json:"weight"`
			AverageDailySteps        int64     `json:"averageDailySteps"`
			StrideLengthWalking      float64   `json:"strideLengthWalking"`
			StrideLengthRunning      float64   `json:"strideLengthRunning"`
			MemberSince              string    `json:"memberSince"`
			AboutMe                  string    `json:"aboutMe"`
			Avatar                   string    `json:"avatar"`
			Avatar150                string    `json:"avatar150"`
			Avatar640                string    `json:"avatar640"`
			TopBadges                []Badge   `json:"topBadges"`
			Features                 *Features `json:"features"`
			Ambassador               bool      `json:"ambassador"`
			ChallengesBeta           bool      `json:"challengesBeta"`
			Corporate                bool      `json:"corporate"`
			CorporateAdmin           bool      `json:"corporateAdmin"`
			IsChild                  bool      `json:"isChild"`
			IsCoach                  bool      `json:"isCoach"`
			AutoStrideEnabled        bool      `json:"autoStrideEnabled"`
			LegalTermsAcceptRequired bool      `json:"legalTermsAcceptRequired"`
			IsBugReportEnabled       bool      `json:"isBugReportEnabled"`
			MfaEnabled               bool      `json:"mfaEnabled"`
			SdkDeveloper             bool      `json:"sdkDeveloper"`
			SleepTracking            string    `json:"sleepTracking"`
			StrideLengthWalkingType  string    `json:"strideLengthWalkingType"`
			StrideLengthRunningType  string    `json:"strideLengthRunningType"`
			ClockTimeDisplayFormat   string    `json:"clockTimeDisplayFormat"`
			StartDayOfWeek           string    `json:"startDayOfWeek"`
			Locale                   string    `json:"locale"`
			LanguageLocale           string    `json:"languageLocale"`
			FoodsLocale              string    `json:"foodsLocale"`
			HeightUnit               string    `json:"heightUnit"`
			WeightUnit               string    `json:"weightUnit"`
			DistanceUnit             string    `json:"distanceUnit"`
			TemperatureUnit          string    `json:"temperatureUnit"`
			SwimUnit                 string    `json:"swimUnit"`
			GlucoseUnit              string    `json:"glucoseUnit"`
			WaterUnit                string    `json:"waterUnit"`
			WaterUnitName            string    `json:"waterUnitName"`
		} `json:"user"`
	}

	// Profile represents user's profile.
	Profile struct {
		EncodedID                string
		DisplayName              string
		DisplayNameSetting       string
		FullName                 string
		FirstName                string
		LastName                 string
		Gender                   string
		DateOfBirth              *time.Time
		Age                      int64
		PhoneNumber              string
		Country                  string
		State                    string
		City                     string
		Timezone                 *time.Location
		OffsetFromUTCMillis      int64
		Height                   float64
		Weight                   float64
		AverageDailySteps        int64
		StrideLengthWalking      float64
		StrideLengthRunning      float64
		MemberSince              *time.Time // MemberSince is the date when the a user registered
		AboutMe                  string
		Avatar                   *url.URL
		Avatar150                *url.URL
		Avatar640                *url.URL
		TopBadges                []Badge
		Features                 *Features
		Ambassador               bool
		ChallengesBeta           bool
		Corporate                bool
		CorporateAdmin           bool
		IsChild                  bool
		IsCoach                  bool
		AutoStrideEnabled        bool
		LegalTermsAcceptRequired bool
		IsBugReportEnabled       bool
		MfaEnabled               bool
		SdkDeveloper             bool
		SleepTracking            string
		StrideLengthWalkingType  string
		StrideLengthRunningType  string
		ClockTimeDisplayFormat   string
		StartDayOfWeek           string
		Locale                   string
		LanguageLocale           string
		FoodsLocale              string
		HeightUnit               string
		WeightUnit               string
		DistanceUnit             string
		TemperatureUnit          string
		SwimUnit                 string
		GlucoseUnit              string
		WaterUnit                string
		WaterUnitName            string
	}
)

// UnmarshalJSON implements the json.Unmarshaler interface.
func (bg *Badge) UnmarshalJSON(b []byte) error {
	var raw rawBadge
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}

	dateTime, err := parseTime("2006-01-02", raw.DateTime)
	if err != nil {
		return err
	}
	image50Px, err := url.Parse(raw.Image50Px)
	if err != nil {
		return err
	}
	image75Px, err := url.Parse(raw.Image75Px)
	if err != nil {
		return err
	}
	image100Px, err := url.Parse(raw.Image100Px)
	if err != nil {
		return err
	}
	image125Px, err := url.Parse(raw.Image125Px)
	if err != nil {
		return err
	}
	image300Px, err := url.Parse(raw.Image300Px)
	if err != nil {
		return err
	}
	shareImage640Px, err := url.Parse(raw.ShareImage640Px)
	if err != nil {
		return err
	}

	bg.EncodedID = raw.EncodedID
	bg.Name = raw.Name
	bg.ShortName = raw.ShortName
	bg.Description = raw.Description
	bg.ShortDescription = raw.ShortDescription
	bg.MobileDescription = raw.MobileDescription
	bg.MarketingDescription = raw.MarketingDescription
	bg.EarnedMessage = raw.EarnedMessage
	bg.Value = raw.Value
	bg.Unit = raw.Unit
	bg.DateTime = dateTime
	bg.TimesAchieved = raw.TimesAchieved
	bg.Cheers = raw.Cheers
	bg.Category = raw.Category
	bg.BadgeType = raw.BadgeType
	bg.BadgeGradientEndColor = raw.BadgeGradientEndColor
	bg.BadgeGradientStartColor = raw.BadgeGradientStartColor
	bg.Image50Px = image50Px
	bg.Image75Px = image75Px
	bg.Image100Px = image100Px
	bg.Image125Px = image125Px
	bg.Image300Px = image300Px
	bg.ShareImage640Px = shareImage640Px
	bg.ShareText = raw.ShareText
	return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (p *Profile) UnmarshalJSON(b []byte) error {
	var raw rawProfile
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}

	dateOfBirth, err := parseTime("2006-01-02", raw.User.DateOfBirth)
	if err != nil {
		return err
	}
	memberSince, err := parseTime("2006-01-02", raw.User.MemberSince)
	if err != nil {
		return err
	}
	timezone, err := time.LoadLocation(raw.User.Timezone)
	if err != nil {
		timezone = time.FixedZone(raw.User.Timezone, int(raw.User.OffsetFromUTCMillis/1e3))
	}
	avatar, err := url.Parse(raw.User.Avatar)
	if err != nil {
		return err
	}
	avatar150, err := url.Parse(raw.User.Avatar150)
	if err != nil {
		return err
	}
	avatar640, err := url.Parse(raw.User.Avatar640)
	if err != nil {
		return err
	}

	p.EncodedID = raw.User.EncodedID
	p.DisplayName = raw.User.DisplayName
	p.DisplayNameSetting = raw.User.DisplayNameSetting
	p.FullName = raw.User.FullName
	p.FirstName = raw.User.FirstName
	p.LastName = raw.User.LastName
	p.Gender = raw.User.Gender
	p.DateOfBirth = dateOfBirth
	p.Age = raw.User.Age
	p.PhoneNumber = raw.User.PhoneNumber
	p.Country = raw.User.Country
	p.State = raw.User.State
	p.City = raw.User.City
	p.Timezone = timezone
	p.OffsetFromUTCMillis = raw.User.OffsetFromUTCMillis
	p.Height = raw.User.Height
	p.Weight = raw.User.Weight
	p.AverageDailySteps = raw.User.AverageDailySteps
	p.StrideLengthWalking = raw.User.StrideLengthWalking
	p.StrideLengthRunning = raw.User.StrideLengthRunning
	p.MemberSince = memberSince
	p.AboutMe = raw.User.AboutMe
	p.Avatar = avatar
	p.Avatar150 = avatar150
	p.Avatar640 = avatar640
	p.TopBadges = raw.User.TopBadges
	p.Features = raw.User.Features
	p.Ambassador = raw.User.Ambassador
	p.ChallengesBeta = raw.User.ChallengesBeta
	p.Corporate = raw.User.Corporate
	p.CorporateAdmin = raw.User.CorporateAdmin
	p.IsChild = raw.User.IsChild
	p.IsCoach = raw.User.IsCoach
	p.AutoStrideEnabled = raw.User.AutoStrideEnabled
	p.LegalTermsAcceptRequired = raw.User.LegalTermsAcceptRequired
	p.IsBugReportEnabled = raw.User.IsBugReportEnabled
	p.MfaEnabled = raw.User.MfaEnabled
	p.SdkDeveloper = raw.User.SdkDeveloper
	p.SleepTracking = raw.User.SleepTracking
	p.StrideLengthWalkingType = raw.User.StrideLengthWalkingType
	p.StrideLengthRunningType = raw.User.StrideLengthRunningType
	p.ClockTimeDisplayFormat = raw.User.ClockTimeDisplayFormat
	p.StartDayOfWeek = raw.User.StartDayOfWeek
	p.Locale = raw.User.Locale
	p.LanguageLocale = raw.User.LanguageLocale
	p.FoodsLocale = raw.User.FoodsLocale
	p.HeightUnit = raw.User.HeightUnit
	p.WeightUnit = raw.User.WeightUnit
	p.DistanceUnit = raw.User.DistanceUnit
	p.TemperatureUnit = raw.User.TemperatureUnit
	p.SwimUnit = raw.User.SwimUnit
	p.GlucoseUnit = raw.User.GlucoseUnit
	p.WaterUnit = raw.User.WaterUnit
	p.WaterUnitName = raw.User.WaterUnitName
	return nil
}

// GetProfile retrieves the user's profile data.
//
// Scope.Profile is required.
//
// Scope.Location and Scope.Nutrition is required to obtain some fields.
//
// Web API Reference: https://dev.fitbit.com/build/reference/web-api/user/get-profile/
func (c *Client) GetProfile(ctx context.Context, userID string, token *Token) (*Profile, *RateLimit, []byte, error) {
	endpoint := c.getEndpoint("GetProfile", userID)
	b, rateLimit, err := c.getRequest(ctx, token, endpoint)
	if err != nil {
		return nil, nil, b, err
	}
	var profile Profile
	if err := json.Unmarshal(b, &profile); err != nil {
		return nil, rateLimit, b, err
	}
	return &profile, rateLimit, b, nil
}
