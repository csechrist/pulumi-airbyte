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
    public sealed class SourceMailchimpConfigurationCredentialsOAuth20
    {
        /// <summary>
        /// An access token generated using the above client ID and secret.
        /// </summary>
        public readonly string AccessToken;
        /// <summary>
        /// The Client ID of your OAuth application.
        /// </summary>
        public readonly string? ClientId;
        /// <summary>
        /// The Client Secret of your OAuth application.
        /// </summary>
        public readonly string? ClientSecret;

        [OutputConstructor]
        private SourceMailchimpConfigurationCredentialsOAuth20(
            string accessToken,

            string? clientId,

            string? clientSecret)
        {
            AccessToken = accessToken;
            ClientId = clientId;
            ClientSecret = clientSecret;
        }
    }
}
