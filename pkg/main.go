package pkg

import (
	s3bucketv1 "buf.build/gen/go/plantoncloud/project-planton/protocolbuffers/go/project/planton/provider/aws/s3bucket/v1"
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func Resources(ctx *pulumi.Context, stackInput *s3bucketv1.S3BucketStackInput) error {
	awsCredential := stackInput.AwsCredential

	//create aws provider using the credentials from the input
	_, err := aws.NewProvider(ctx,
		"native-provider",
		&aws.ProviderArgs{
			AccessKey: pulumi.String(awsCredential.AccessKeyId),
			SecretKey: pulumi.String(awsCredential.SecretAccessKey),
			Region:    pulumi.String(awsCredential.Region),
		})
	if err != nil {
		return errors.Wrap(err, "failed to create aws provider")
	}

	return nil
}
