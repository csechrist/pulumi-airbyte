// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class SourceGoogleAnalyticsDataApiConfigurationCustomReportsArrayMetricFilterOrGroupExpressionFilterNumericFilterValueDoubleValueArgs : global::Pulumi.ResourceArgs
    {
        [Input("value", required: true)]
        public Input<double> Value { get; set; } = null!;

        public SourceGoogleAnalyticsDataApiConfigurationCustomReportsArrayMetricFilterOrGroupExpressionFilterNumericFilterValueDoubleValueArgs()
        {
        }
        public static new SourceGoogleAnalyticsDataApiConfigurationCustomReportsArrayMetricFilterOrGroupExpressionFilterNumericFilterValueDoubleValueArgs Empty => new SourceGoogleAnalyticsDataApiConfigurationCustomReportsArrayMetricFilterOrGroupExpressionFilterNumericFilterValueDoubleValueArgs();
    }
}
