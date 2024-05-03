// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class SourceAwsCloudtrailConfigurationGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("awsKeyId", required: true)]
        private Input<string>? _awsKeyId;

        /// <summary>
        /// AWS CloudTrail Access Key ID. See the \n\ndocs\n\n for more information on how to obtain this key.
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

        /// <summary>
        /// The default AWS Region to use, for example, us-west-1 or us-west-2. When specifying a Region inline during client initialization, this property is named region_name.
        /// </summary>
        [Input("awsRegionName", required: true)]
        public Input<string> AwsRegionName { get; set; } = null!;

        [Input("awsSecretKey", required: true)]
        private Input<string>? _awsSecretKey;

        /// <summary>
        /// AWS CloudTrail Access Key ID. See the \n\ndocs\n\n for more information on how to obtain this key.
        /// </summary>
        public Input<string>? AwsSecretKey
        {
            get => _awsSecretKey;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _awsSecretKey = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// The date you would like to replicate data. Data in AWS CloudTrail is available for last 90 days only. Format: YYYY-MM-DD. Default: "1970-01-01"
        /// </summary>
        [Input("startDate")]
        public Input<string>? StartDate { get; set; }

        public SourceAwsCloudtrailConfigurationGetArgs()
        {
        }
        public static new SourceAwsCloudtrailConfigurationGetArgs Empty => new SourceAwsCloudtrailConfigurationGetArgs();
    }
}