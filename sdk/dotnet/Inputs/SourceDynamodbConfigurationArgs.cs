// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class SourceDynamodbConfigurationArgs : global::Pulumi.ResourceArgs
    {
        [Input("accessKeyId", required: true)]
        private Input<string>? _accessKeyId;

        /// <summary>
        /// The access key id to access Dynamodb. Airbyte requires read permissions to the database
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
        /// the URL of the Dynamodb database. Default: ""
        /// </summary>
        [Input("endpoint")]
        public Input<string>? Endpoint { get; set; }

        /// <summary>
        /// The region of the Dynamodb database. must be one of ["", "af-south-1", "ap-east-1", "ap-northeast-1", "ap-northeast-2", "ap-northeast-3", "ap-south-1", "ap-south-2", "ap-southeast-1", "ap-southeast-2", "ap-southeast-3", "ap-southeast-4", "ca-central-1", "ca-west-1", "cn-north-1", "cn-northwest-1", "eu-central-1", "eu-central-2", "eu-north-1", "eu-south-1", "eu-south-2", "eu-west-1", "eu-west-2", "eu-west-3", "il-central-1", "me-central-1", "me-south-1", "sa-east-1", "us-east-1", "us-east-2", "us-gov-east-1", "us-gov-west-1", "us-west-1", "us-west-2"]; Default: ""
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// Comma separated reserved attribute names present in your tables
        /// </summary>
        [Input("reservedAttributeNames")]
        public Input<string>? ReservedAttributeNames { get; set; }

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

        public SourceDynamodbConfigurationArgs()
        {
        }
        public static new SourceDynamodbConfigurationArgs Empty => new SourceDynamodbConfigurationArgs();
    }
}