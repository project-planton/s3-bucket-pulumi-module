package outputs

import (
	"github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/aws/s3bucket"
	"github.com/pulumi/pulumi/sdk/v3/go/auto"
)

func PulumiOutputsToStackOutputsConverter(pulumiOutputs auto.OutputMap,
	input *s3bucket.S3BucketStackInput) *s3bucket.S3BucketStackOutputs {
	return &s3bucket.S3BucketStackOutputs{}
}
