// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class SourceGoogleAnalyticsDataApiConfigurationCustomReportsArrayMetricFilterOrGroupExpressionFilterNumericFilterGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("operations", required: true)]
        private InputList<string>? _operations;
        public InputList<string> Operations
        {
            get => _operations ?? (_operations = new InputList<string>());
            set => _operations = value;
        }

        [Input("value", required: true)]
        public Input<Inputs.SourceGoogleAnalyticsDataApiConfigurationCustomReportsArrayMetricFilterOrGroupExpressionFilterNumericFilterValueGetArgs> Value { get; set; } = null!;

        public SourceGoogleAnalyticsDataApiConfigurationCustomReportsArrayMetricFilterOrGroupExpressionFilterNumericFilterGetArgs()
        {
        }
        public static new SourceGoogleAnalyticsDataApiConfigurationCustomReportsArrayMetricFilterOrGroupExpressionFilterNumericFilterGetArgs Empty => new SourceGoogleAnalyticsDataApiConfigurationCustomReportsArrayMetricFilterOrGroupExpressionFilterNumericFilterGetArgs();
    }
}