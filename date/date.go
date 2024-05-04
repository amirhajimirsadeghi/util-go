package date

import (
	"fmt"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

// Introducing a special type to help with handling dates.
// Especially when dealing with marshalling and unmarshalling
type Date time.Time

func (d Date) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("\"%s\"", time.Time(d).Format("2006-01-02"))
	return []byte(stamp), nil
}

func (d *Date) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")

	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}

	*d = Date(t)
	return nil
}

func (d Date) MarshalDynamoDBAttributeValue() (types.AttributeValue, error) {
	return &types.AttributeValueMemberS{Value: time.Time(d).Format("2006-01-02")}, nil
}

func (p *Date) UnmarshalDynamoDBAttributeValue(av types.AttributeValue) error {
	s, ok := av.(*types.AttributeValueMemberS)
	if !ok {
		return fmt.Errorf("attribute value is not a string")
	}

	t, err := time.Parse("2006-01-02", s.Value)
	if err != nil {
		return err
	}

	*p = Date(t)
	return nil
}
