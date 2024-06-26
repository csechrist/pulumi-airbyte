// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class SourceGoogleAnalyticsDataApiConfigurationCustomReportsArrayCohortSpecEnabledCohortArgs : global::Pulumi.ResourceArgs
    {
        [Input("dateRange", required: true)]
        public Input<Inputs.SourceGoogleAnalyticsDataApiConfigurationCustomReportsArrayCohortSpecEnabledCohortDateRangeArgs> DateRange { get; set; } = null!;

        /// <summary>
        /// Dimension used by the cohort. Required and only supports `firstSessionDate`. must be one of ["firstSessionDate"]
        /// </summary>
        [Input("dimension", required: true)]
        public Input<string> Dimension { get; set; } = null!;

        /// <summary>
        /// Assigns a name to this cohort. If not set, cohorts are named by their zero based index cohort*0, cohort*1, etc.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public SourceGoogleAnalyticsDataApiConfigurationCustomReportsArrayCohortSpecEnabledCohortArgs()
        {
        }
        public static new SourceGoogleAnalyticsDataApiConfigurationCustomReportsArrayCohortSpecEnabledCohortArgs Empty => new SourceGoogleAnalyticsDataApiConfigurationCustomReportsArrayCohortSpecEnabledCohortArgs();
    }
}
