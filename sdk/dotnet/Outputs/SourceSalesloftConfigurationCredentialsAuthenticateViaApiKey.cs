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
    public sealed class SourceSalesloftConfigurationCredentialsAuthenticateViaApiKey
    {
        /// <summary>
        /// API Key for making authenticated requests. More instruction on how to find this value in our \n\ndocs\n\n
        /// </summary>
        public readonly string ApiKey;

        [OutputConstructor]
        private SourceSalesloftConfigurationCredentialsAuthenticateViaApiKey(string apiKey)
        {
            ApiKey = apiKey;
        }
    }
}