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
    public sealed class SourceGridlyConfiguration
    {
        public readonly string ApiKey;
        /// <summary>
        /// ID of a grid, or can be ID of a branch
        /// </summary>
        public readonly string GridId;

        [OutputConstructor]
        private SourceGridlyConfiguration(
            string apiKey,

            string gridId)
        {
            ApiKey = apiKey;
            GridId = gridId;
        }
    }
}