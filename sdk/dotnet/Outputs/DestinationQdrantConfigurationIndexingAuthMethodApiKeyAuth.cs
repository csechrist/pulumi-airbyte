// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Outputs
{

    [OutputType]
    public sealed class DestinationQdrantConfigurationIndexingAuthMethodApiKeyAuth
    {
        /// <summary>
        /// API Key for the Qdrant instance
        /// </summary>
        public readonly string ApiKey;

        [OutputConstructor]
        private DestinationQdrantConfigurationIndexingAuthMethodApiKeyAuth(string apiKey)
        {
            ApiKey = apiKey;
        }
    }
}
