// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class DestinationRedshiftConfigurationUploadingMethodAwss3StagingArgs : global::Pulumi.ResourceArgs
    {
        [Input("accessKeyId", required: true)]
        private Input<string>? _accessKeyId;

        /// <summary>
        /// This ID grants access to the above S3 staging bucket. Airbyte requires Read and Write permissions to the given bucket. See \n\nAWS docs\n\n on how to generate an access key ID and secret access key.
        /// </summary>
        public Input<string>? AccessKeyId
        {
            get => _accessKeyId;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _accessKeyId = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// How to encrypt the staging data
        /// </summary>
        [Input("encryption")]
        public Input<Inputs.DestinationRedshiftConfigurationUploadingMethodAwss3StagingEncryptionArgs>? Encryption { get; set; }

        /// <summary>
        /// Number of file buffers allocated for writing data. Increasing this number is beneficial for connections using Change Data Capture (CDC) and up to the number of streams within a connection. Increasing the number of file buffers past the maximum number of streams has deteriorating effects. Default: 10
        /// </summary>
        [Input("fileBufferCount")]
        public Input<int>? FileBufferCount { get; set; }

        /// <summary>
        /// The pattern allows you to set the file-name format for the S3 staging file(s)
        /// </summary>
        [Input("fileNamePattern")]
        public Input<string>? FileNamePattern { get; set; }

        /// <summary>
        /// Whether to delete the staging files from S3 after completing the sync. See \n\n docs\n\n for details. Default: true
        /// </summary>
        [Input("purgeStagingData")]
        public Input<bool>? PurgeStagingData { get; set; }

        /// <summary>
        /// The name of the staging S3 bucket.
        /// </summary>
        [Input("s3BucketName", required: true)]
        public Input<string> S3BucketName { get; set; } = null!;

        /// <summary>
        /// The directory under the S3 bucket where data will be written. If not provided, then defaults to the root directory. See \n\npath's name recommendations\n\n for more details.
        /// </summary>
        [Input("s3BucketPath")]
        public Input<string>? S3BucketPath { get; set; }

        /// <summary>
        /// The region of the S3 staging bucket. must be one of ["", "af-south-1", "ap-east-1", "ap-northeast-1", "ap-northeast-2", "ap-northeast-3", "ap-south-1", "ap-south-2", "ap-southeast-1", "ap-southeast-2", "ap-southeast-3", "ap-southeast-4", "ca-central-1", "ca-west-1", "cn-north-1", "cn-northwest-1", "eu-central-1", "eu-central-2", "eu-north-1", "eu-south-1", "eu-south-2", "eu-west-1", "eu-west-2", "eu-west-3", "il-central-1", "me-central-1", "me-south-1", "sa-east-1", "us-east-1", "us-east-2", "us-gov-east-1", "us-gov-west-1", "us-west-1", "us-west-2"]; Default: ""
        /// </summary>
        [Input("s3BucketRegion")]
        public Input<string>? S3BucketRegion { get; set; }

        [Input("secretAccessKey", required: true)]
        private Input<string>? _secretAccessKey;

        /// <summary>
        /// The corresponding secret to the above access key id. See \n\nAWS docs\n\n on how to generate an access key ID and secret access key.
        /// </summary>
        public Input<string>? SecretAccessKey
        {
            get => _secretAccessKey;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _secretAccessKey = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        public DestinationRedshiftConfigurationUploadingMethodAwss3StagingArgs()
        {
        }
        public static new DestinationRedshiftConfigurationUploadingMethodAwss3StagingArgs Empty => new DestinationRedshiftConfigurationUploadingMethodAwss3StagingArgs();
    }
}
