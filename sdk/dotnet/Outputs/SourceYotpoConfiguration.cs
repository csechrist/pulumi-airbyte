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
    public sealed class SourceYotpoConfiguration
    {
        /// <summary>
        /// Access token recieved as a result of API call to https://api.yotpo.com/oauth/token (Ref- https://apidocs.yotpo.com/reference/yotpo-authentication)
        /// </summary>
        public readonly string AccessToken;
        /// <summary>
        /// App key found at settings (Ref- https://settings.yotpo.com/#/general_settings)
        /// </summary>
        public readonly string AppKey;
        /// <summary>
        /// Email address registered with yotpo. Default: "example@gmail.com"
        /// </summary>
        public readonly string? Email;
        /// <summary>
        /// Date time filter for incremental filter, Specify which date to extract from.
        /// </summary>
        public readonly string StartDate;

        [OutputConstructor]
        private SourceYotpoConfiguration(
            string accessToken,

            string appKey,

            string? email,

            string startDate)
        {
            AccessToken = accessToken;
            AppKey = appKey;
            Email = email;
            StartDate = startDate;
        }
    }
}