// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class SourceSurveySparrowConfigurationRegionArgs : global::Pulumi.ResourceArgs
    {
        [Input("euBasedAccount")]
        public Input<Inputs.SourceSurveySparrowConfigurationRegionEuBasedAccountArgs>? EuBasedAccount { get; set; }

        [Input("globalAccount")]
        public Input<Inputs.SourceSurveySparrowConfigurationRegionGlobalAccountArgs>? GlobalAccount { get; set; }

        public SourceSurveySparrowConfigurationRegionArgs()
        {
        }
        public static new SourceSurveySparrowConfigurationRegionArgs Empty => new SourceSurveySparrowConfigurationRegionArgs();
    }
}