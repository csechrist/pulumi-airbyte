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
    public sealed class SourceGoogleAnalyticsDataApiConfigurationCustomReportsArrayMetricFilterOrGroupExpressionFilterBetweenFilterFromValue
    {
        public readonly Outputs.SourceGoogleAnalyticsDataApiConfigurationCustomReportsArrayMetricFilterOrGroupExpressionFilterBetweenFilterFromValueDoubleValue? DoubleValue;
        public readonly Outputs.SourceGoogleAnalyticsDataApiConfigurationCustomReportsArrayMetricFilterOrGroupExpressionFilterBetweenFilterFromValueInt64Value? Int64Value;

        [OutputConstructor]
        private SourceGoogleAnalyticsDataApiConfigurationCustomReportsArrayMetricFilterOrGroupExpressionFilterBetweenFilterFromValue(
            Outputs.SourceGoogleAnalyticsDataApiConfigurationCustomReportsArrayMetricFilterOrGroupExpressionFilterBetweenFilterFromValueDoubleValue? doubleValue,

            Outputs.SourceGoogleAnalyticsDataApiConfigurationCustomReportsArrayMetricFilterOrGroupExpressionFilterBetweenFilterFromValueInt64Value? int64Value)
        {
            DoubleValue = doubleValue;
            Int64Value = int64Value;
        }
    }
}