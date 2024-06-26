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
    public sealed class SourceSalesloftConfiguration
    {
        public readonly Outputs.SourceSalesloftConfigurationCredentials Credentials;
        /// <summary>
        /// The date from which you'd like to replicate data for Salesloft API, in the format YYYY-MM-DDT00:00:00Z. All data generated after this date will be replicated.
        /// </summary>
        public readonly string StartDate;

        [OutputConstructor]
        private SourceSalesloftConfiguration(
            Outputs.SourceSalesloftConfigurationCredentials credentials,

            string startDate)
        {
            Credentials = credentials;
            StartDate = startDate;
        }
    }
}
