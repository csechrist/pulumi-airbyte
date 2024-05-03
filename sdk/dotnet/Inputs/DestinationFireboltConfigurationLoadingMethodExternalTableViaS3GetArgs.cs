// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class DestinationFireboltConfigurationLoadingMethodExternalTableViaS3GetArgs : global::Pulumi.ResourceArgs
    {
        [Input("awsKeyId", required: true)]
        private Input<string>? _awsKeyId;

        /// <summary>
        /// AWS access key granting read and write access to S3.
        /// </summary>
        public Input<string>? AwsKeyId
        {
            get => _awsKeyId;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _awsKeyId = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("awsKeySecret", required: true)]
        private Input<string>? _awsKeySecret;

        /// <summary>
        /// Corresponding secret part of the AWS Key
        /// </summary>
        public Input<string>? AwsKeySecret
        {
            get => _awsKeySecret;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _awsKeySecret = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// The name of the S3 bucket.
        /// </summary>
        [Input("s3Bucket", required: true)]
        public Input<string> S3Bucket { get; set; } = null!;

        /// <summary>
        /// Region name of the S3 bucket.
        /// </summary>
        [Input("s3Region", required: true)]
        public Input<string> S3Region { get; set; } = null!;

        public DestinationFireboltConfigurationLoadingMethodExternalTableViaS3GetArgs()
        {
        }
        public static new DestinationFireboltConfigurationLoadingMethodExternalTableViaS3GetArgs Empty => new DestinationFireboltConfigurationLoadingMethodExternalTableViaS3GetArgs();
    }
}