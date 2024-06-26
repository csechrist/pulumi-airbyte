// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class DestinationDynamodbConfigurationArgs : global::Pulumi.ResourceArgs
    {
        [Input("accessKeyId", required: true)]
        private Input<string>? _accessKeyId;

        /// <summary>
        /// The access key id to access the DynamoDB. Airbyte requires Read and Write permissions to the DynamoDB.
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
        /// This is your DynamoDB endpoint url.(if you are working with AWS DynamoDB, just leave empty). Default: ""
        /// </summary>
        [Input("dynamodbEndpoint")]
        public Input<string>? DynamodbEndpoint { get; set; }

        /// <summary>
        /// The region of the DynamoDB. must be one of ["", "af-south-1", "ap-east-1", "ap-northeast-1", "ap-northeast-2", "ap-northeast-3", "ap-south-1", "ap-south-2", "ap-southeast-1", "ap-southeast-2", "ap-southeast-3", "ap-southeast-4", "ca-central-1", "ca-west-1", "cn-north-1", "cn-northwest-1", "eu-central-1", "eu-central-2", "eu-north-1", "eu-south-1", "eu-south-2", "eu-west-1", "eu-west-2", "eu-west-3", "il-central-1", "me-central-1", "me-south-1", "sa-east-1", "us-east-1", "us-east-2", "us-gov-east-1", "us-gov-west-1", "us-west-1", "us-west-2"]; Default: ""
        /// </summary>
        [Input("dynamodbRegion")]
        public Input<string>? DynamodbRegion { get; set; }

        /// <summary>
        /// The prefix to use when naming DynamoDB tables.
        /// </summary>
        [Input("dynamodbTableNamePrefix", required: true)]
        public Input<string> DynamodbTableNamePrefix { get; set; } = null!;

        [Input("secretAccessKey", required: true)]
        private Input<string>? _secretAccessKey;

        /// <summary>
        /// The corresponding secret to the access key id.
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

        public DestinationDynamodbConfigurationArgs()
        {
        }
        public static new DestinationDynamodbConfigurationArgs Empty => new DestinationDynamodbConfigurationArgs();
    }
}
