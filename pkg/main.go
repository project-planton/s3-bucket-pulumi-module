package pkg

import (
	"github.com/pkg/errors"
	"github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/aws/s3bucket"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ResourceStack struct {
	StackInput *s3bucket.S3BucketStackInput
}

func (s *ResourceStack) Resources(ctx *pulumi.Context) error {
	awsCredential := s.StackInput.AwsCredential

	//create aws provider using the credentials from the input
	_, err := aws.NewProvider(ctx,
		"native-provider",
		&aws.ProviderArgs{
			AccessKey: pulumi.String(awsCredential.Spec.AccessKeyId),
			SecretKey: pulumi.String(awsCredential.Spec.SecretAccessKey),
			Region:    pulumi.String(awsCredential.Spec.Region),
		})
	if err != nil {
		return errors.Wrap(err, "failed to create aws provider")
	}

	return nil
}
