// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class SourceGoogleAnalyticsDataApiConfigurationCustomReportsArrayDimensionFilterOrGroupExpressionFilterArgs : global::Pulumi.ResourceArgs
    {
        [Input("betweenFilter")]
        public Input<Inputs.SourceGoogleAnalyticsDataApiConfigurationCustomReportsArrayDimensionFilterOrGroupExpressionFilterBetweenFilterArgs>? BetweenFilter { get; set; }

        [Input("inListFilter")]
        public Input<Inputs.SourceGoogleAnalyticsDataApiConfigurationCustomReportsArrayDimensionFilterOrGroupExpressionFilterInListFilterArgs>? InListFilter { get; set; }

        [Input("numericFilter")]
        public Input<Inputs.SourceGoogleAnalyticsDataApiConfigurationCustomReportsArrayDimensionFilterOrGroupExpressionFilterNumericFilterArgs>? NumericFilter { get; set; }

        [Input("stringFilter")]
        public Input<Inputs.SourceGoogleAnalyticsDataApiConfigurationCustomReportsArrayDimensionFilterOrGroupExpressionFilterStringFilterArgs>? StringFilter { get; set; }

        public SourceGoogleAnalyticsDataApiConfigurationCustomReportsArrayDimensionFilterOrGroupExpressionFilterArgs()
        {
        }
        public static new SourceGoogleAnalyticsDataApiConfigurationCustomReportsArrayDimensionFilterOrGroupExpressionFilterArgs Empty => new SourceGoogleAnalyticsDataApiConfigurationCustomReportsArrayDimensionFilterOrGroupExpressionFilterArgs();
    }
}