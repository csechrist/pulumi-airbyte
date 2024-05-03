// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class SourceGoogleAnalyticsDataApiConfigurationCustomReportsArrayCohortSpecEnabledGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Optional settings for a cohort report.
        /// </summary>
        [Input("cohortReportSettings")]
        public Input<Inputs.SourceGoogleAnalyticsDataApiConfigurationCustomReportsArrayCohortSpecEnabledCohortReportSettingsGetArgs>? CohortReportSettings { get; set; }

        [Input("cohorts")]
        private InputList<Inputs.SourceGoogleAnalyticsDataApiConfigurationCustomReportsArrayCohortSpecEnabledCohortGetArgs>? _cohorts;
        public InputList<Inputs.SourceGoogleAnalyticsDataApiConfigurationCustomReportsArrayCohortSpecEnabledCohortGetArgs> Cohorts
        {
            get => _cohorts ?? (_cohorts = new InputList<Inputs.SourceGoogleAnalyticsDataApiConfigurationCustomReportsArrayCohortSpecEnabledCohortGetArgs>());
            set => _cohorts = value;
        }

        [Input("cohortsRange")]
        public Input<Inputs.SourceGoogleAnalyticsDataApiConfigurationCustomReportsArrayCohortSpecEnabledCohortsRangeGetArgs>? CohortsRange { get; set; }

        public SourceGoogleAnalyticsDataApiConfigurationCustomReportsArrayCohortSpecEnabledGetArgs()
        {
        }
        public static new SourceGoogleAnalyticsDataApiConfigurationCustomReportsArrayCohortSpecEnabledGetArgs Empty => new SourceGoogleAnalyticsDataApiConfigurationCustomReportsArrayCohortSpecEnabledGetArgs();
    }
}