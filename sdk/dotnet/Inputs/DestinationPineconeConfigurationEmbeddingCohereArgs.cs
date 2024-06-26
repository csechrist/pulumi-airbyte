// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class DestinationPineconeConfigurationEmbeddingCohereArgs : global::Pulumi.ResourceArgs
    {
        [Input("cohereKey", required: true)]
        private Input<string>? _cohereKey;
        public Input<string>? CohereKey
        {
            get => _cohereKey;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _cohereKey = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        public DestinationPineconeConfigurationEmbeddingCohereArgs()
        {
        }
        public static new DestinationPineconeConfigurationEmbeddingCohereArgs Empty => new DestinationPineconeConfigurationEmbeddingCohereArgs();
    }
}
