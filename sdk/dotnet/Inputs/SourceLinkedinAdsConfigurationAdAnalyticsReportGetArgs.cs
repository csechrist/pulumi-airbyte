// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class SourceLinkedinAdsConfigurationAdAnalyticsReportGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name for the custom report.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Choose a category to pivot your analytics report around. This selection will organize your data based on the chosen attribute, allowing you to analyze trends and performance from different perspectives. must be one of ["COMPANY", "ACCOUNT", "SHARE", "CAMPAIGN", "CREATIVE", "CAMPAIGN*GROUP", "CONVERSION", "CONVERSATION*NODE", "CONVERSATION*NODE*OPTION*INDEX", "SERVING*LOCATION", "CARD*INDEX", "MEMBER*COMPANY*SIZE", "MEMBER*INDUSTRY", "MEMBER*SENIORITY", "MEMBER*JOB*TITLE", "MEMBER*JOB*FUNCTION", "MEMBER*COUNTRY*V2", "MEMBER*REGION*V2", "MEMBER*COMPANY", "PLACEMENT*NAME", "IMPRESSION*DEVICE_TYPE"]
        /// </summary>
        [Input("pivotBy", required: true)]
        public Input<string> PivotBy { get; set; } = null!;

        /// <summary>
        /// Choose how to group the data in your report by time. The options are:\n\n- 'ALL': A single result summarizing the entire time range.\n\n- 'DAILY': Group results by each day.\n\n- 'MONTHLY': Group results by each month.\n\n- 'YEARLY': Group results by each year.\n\nSelecting a time grouping helps you analyze trends and patterns over different time periods. must be one of ["ALL", "DAILY", "MONTHLY", "YEARLY"]
        /// </summary>
        [Input("timeGranularity", required: true)]
        public Input<string> TimeGranularity { get; set; } = null!;

        public SourceLinkedinAdsConfigurationAdAnalyticsReportGetArgs()
        {
        }
        public static new SourceLinkedinAdsConfigurationAdAnalyticsReportGetArgs Empty => new SourceLinkedinAdsConfigurationAdAnalyticsReportGetArgs();
    }
}