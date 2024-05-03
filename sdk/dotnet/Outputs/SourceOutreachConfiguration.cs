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
    public sealed class SourceOutreachConfiguration
    {
        /// <summary>
        /// The Client ID of your Outreach developer application.
        /// </summary>
        public readonly string ClientId;
        /// <summary>
        /// The Client Secret of your Outreach developer application.
        /// </summary>
        public readonly string ClientSecret;
        /// <summary>
        /// A Redirect URI is the location where the authorization server sends the user once the app has been successfully authorized and granted an authorization code or access token.
        /// </summary>
        public readonly string RedirectUri;
        /// <summary>
        /// The token for obtaining the new access token.
        /// </summary>
        public readonly string RefreshToken;
        /// <summary>
        /// The date from which you'd like to replicate data for Outreach API, in the format YYYY-MM-DDT00:00:00Z. All data generated after this date will be replicated.
        /// </summary>
        public readonly string StartDate;

        [OutputConstructor]
        private SourceOutreachConfiguration(
            string clientId,

            string clientSecret,

            string redirectUri,

            string refreshToken,

            string startDate)
        {
            ClientId = clientId;
            ClientSecret = clientSecret;
            RedirectUri = redirectUri;
            RefreshToken = refreshToken;
            StartDate = startDate;
        }
    }
}