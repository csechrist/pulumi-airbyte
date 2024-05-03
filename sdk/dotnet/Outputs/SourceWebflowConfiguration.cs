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
    public sealed class SourceWebflowConfiguration
    {
        /// <summary>
        /// The version of the Webflow API to use. See https://developers.webflow.com/#versioning
        /// </summary>
        public readonly string? AcceptVersion;
        /// <summary>
        /// The API token for authenticating to Webflow. See https://university.webflow.com/lesson/intro-to-the-webflow-api
        /// </summary>
        public readonly string ApiKey;
        /// <summary>
        /// The id of the Webflow site you are requesting data from. See https://developers.webflow.com/#sites
        /// </summary>
        public readonly string SiteId;

        [OutputConstructor]
        private SourceWebflowConfiguration(
            string? acceptVersion,

            string apiKey,

            string siteId)
        {
            AcceptVersion = acceptVersion;
            ApiKey = apiKey;
            SiteId = siteId;
        }
    }
}