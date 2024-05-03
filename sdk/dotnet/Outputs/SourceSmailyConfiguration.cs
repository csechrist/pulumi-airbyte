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
    public sealed class SourceSmailyConfiguration
    {
        /// <summary>
        /// API user password. See https://smaily.com/help/api/general/create-api-user/
        /// </summary>
        public readonly string ApiPassword;
        /// <summary>
        /// API Subdomain. See https://smaily.com/help/api/general/create-api-user/
        /// </summary>
        public readonly string ApiSubdomain;
        /// <summary>
        /// API user username. See https://smaily.com/help/api/general/create-api-user/
        /// </summary>
        public readonly string ApiUsername;

        [OutputConstructor]
        private SourceSmailyConfiguration(
            string apiPassword,

            string apiSubdomain,

            string apiUsername)
        {
            ApiPassword = apiPassword;
            ApiSubdomain = apiSubdomain;
            ApiUsername = apiUsername;
        }
    }
}