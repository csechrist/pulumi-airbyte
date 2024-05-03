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
    public sealed class SourceGoogleWebfontsConfiguration
    {
        /// <summary>
        /// Optional, Available params- json, media, proto
        /// </summary>
        public readonly string? Alt;
        /// <summary>
        /// API key is required to access google apis, For getting your's goto google console and generate api key for Webfonts
        /// </summary>
        public readonly string ApiKey;
        /// <summary>
        /// Optional, boolean type
        /// </summary>
        public readonly string? PrettyPrint;
        /// <summary>
        /// Optional, to find how to sort
        /// </summary>
        public readonly string? Sort;

        [OutputConstructor]
        private SourceGoogleWebfontsConfiguration(
            string? alt,

            string apiKey,

            string? prettyPrint,

            string? sort)
        {
            Alt = alt;
            ApiKey = apiKey;
            PrettyPrint = prettyPrint;
            Sort = sort;
        }
    }
}