// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class SourceGoogleAnalyticsDataApiConfigurationCustomReportsArrayDimensionFilterFilterFilterBetweenFilterToValueGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("doubleValue")]
        public Input<Inputs.SourceGoogleAnalyticsDataApiConfigurationCustomReportsArrayDimensionFilterFilterFilterBetweenFilterToValueDoubleValueGetArgs>? DoubleValue { get; set; }

        [Input("int64Value")]
        public Input<Inputs.SourceGoogleAnalyticsDataApiConfigurationCustomReportsArrayDimensionFilterFilterFilterBetweenFilterToValueInt64ValueGetArgs>? Int64Value { get; set; }

        public SourceGoogleAnalyticsDataApiConfigurationCustomReportsArrayDimensionFilterFilterFilterBetweenFilterToValueGetArgs()
        {
        }
        public static new SourceGoogleAnalyticsDataApiConfigurationCustomReportsArrayDimensionFilterFilterFilterBetweenFilterToValueGetArgs Empty => new SourceGoogleAnalyticsDataApiConfigurationCustomReportsArrayDimensionFilterFilterFilterBetweenFilterToValueGetArgs();
    }
}