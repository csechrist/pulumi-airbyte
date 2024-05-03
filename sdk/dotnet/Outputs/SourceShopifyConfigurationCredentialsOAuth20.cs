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
    public sealed class SourceShopifyConfigurationCredentialsOAuth20
    {
        /// <summary>
        /// The Access Token for making authenticated requests.
        /// </summary>
        public readonly string? AccessToken;
        /// <summary>
        /// The Client ID of the Shopify developer application.
        /// </summary>
        public readonly string? ClientId;
        /// <summary>
        /// The Client Secret of the Shopify developer application.
        /// </summary>
        public readonly string? ClientSecret;

        [OutputConstructor]
        private SourceShopifyConfigurationCredentialsOAuth20(
            string? accessToken,

            string? clientId,

            string? clientSecret)
        {
            AccessToken = accessToken;
            ClientId = clientId;
            ClientSecret = clientSecret;
        }
    }
}