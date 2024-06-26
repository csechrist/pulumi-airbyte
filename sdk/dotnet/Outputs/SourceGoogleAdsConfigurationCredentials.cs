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
    public sealed class SourceGoogleAdsConfigurationCredentials
    {
        /// <summary>
        /// The Access Token for making authenticated requests. For detailed instructions on finding this value, refer to our \n\ndocumentation\n\n.
        /// </summary>
        public readonly string? AccessToken;
        /// <summary>
        /// The Client ID of your Google Ads developer application. For detailed instructions on finding this value, refer to our \n\ndocumentation\n\n.
        /// </summary>
        public readonly string ClientId;
        /// <summary>
        /// The Client Secret of your Google Ads developer application. For detailed instructions on finding this value, refer to our \n\ndocumentation\n\n.
        /// </summary>
        public readonly string ClientSecret;
        /// <summary>
        /// The Developer Token granted by Google to use their APIs. For detailed instructions on finding this value, refer to our \n\ndocumentation\n\n.
        /// </summary>
        public readonly string DeveloperToken;
        /// <summary>
        /// The token used to obtain a new Access Token. For detailed instructions on finding this value, refer to our \n\ndocumentation\n\n.
        /// </summary>
        public readonly string RefreshToken;

        [OutputConstructor]
        private SourceGoogleAdsConfigurationCredentials(
            string? accessToken,

            string clientId,

            string clientSecret,

            string developerToken,

            string refreshToken)
        {
            AccessToken = accessToken;
            ClientId = clientId;
            ClientSecret = clientSecret;
            DeveloperToken = developerToken;
            RefreshToken = refreshToken;
        }
    }
}
