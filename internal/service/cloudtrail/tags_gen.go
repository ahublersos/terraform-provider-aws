// Code generated by internal/generate/tags/main.go; DO NOT EDIT.
package cloudtrail

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudtrail"
	"github.com/aws/aws-sdk-go/service/cloudtrail/cloudtrailiface"
	tftags "github.com/hashicorp/terraform-provider-aws/internal/tags"
)

// ListTags lists cloudtrail service tags.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
func ListTags(conn cloudtrailiface.CloudTrailAPI, identifier string) (tftags.KeyValueTags, error) {
	input := &cloudtrail.ListTagsInput{
		ResourceIdList: aws.StringSlice([]string{identifier}),
	}

	output, err := conn.ListTags(input)

	if err != nil {
		return tftags.New(nil), err
	}

	return KeyValueTags(output.ResourceTagList[0].TagsList), nil
}

// []*SERVICE.Tag handling

// Tags returns cloudtrail service tags.
func Tags(tags tftags.KeyValueTags) []*cloudtrail.Tag {
	result := make([]*cloudtrail.Tag, 0, len(tags))

	for k, v := range tags.Map() {
		tag := &cloudtrail.Tag{
			Key:   aws.String(k),
			Value: aws.String(v),
		}

		result = append(result, tag)
	}

	return result
}

// KeyValueTags creates tftags.KeyValueTags from cloudtrail service tags.
func KeyValueTags(tags []*cloudtrail.Tag) tftags.KeyValueTags {
	m := make(map[string]*string, len(tags))

	for _, tag := range tags {
		m[aws.StringValue(tag.Key)] = tag.Value
	}

	return tftags.New(m)
}

// UpdateTags updates cloudtrail service tags.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
func UpdateTags(conn cloudtrailiface.CloudTrailAPI, identifier string, oldTagsMap interface{}, newTagsMap interface{}) error {
	oldTags := tftags.New(oldTagsMap)
	newTags := tftags.New(newTagsMap)

	if removedTags := oldTags.Removed(newTags); len(removedTags) > 0 {
		input := &cloudtrail.RemoveTagsInput{
			ResourceId: aws.String(identifier),
			TagsList:   Tags(removedTags.IgnoreAWS()),
		}

		_, err := conn.RemoveTags(input)

		if err != nil {
			return fmt.Errorf("error untagging resource (%s): %w", identifier, err)
		}
	}

	if updatedTags := oldTags.Updated(newTags); len(updatedTags) > 0 {
		input := &cloudtrail.AddTagsInput{
			ResourceId: aws.String(identifier),
			TagsList:   Tags(updatedTags.IgnoreAWS()),
		}

		_, err := conn.AddTags(input)

		if err != nil {
			return fmt.Errorf("error tagging resource (%s): %w", identifier, err)
		}
	}

	return nil
}
