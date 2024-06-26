// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class SourceBingAdsConfigurationAccountNameArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Account Name is a string value for comparing with the specified predicate.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// An Operator that will be used to filter accounts. The Contains predicate has features for matching words, matching inflectional forms of words, searching using wildcard characters, and searching using proximity. The Equals is used to return all rows where account name is equal(=) to the string that you provided. must be one of ["Contains", "Equals"]
        /// </summary>
        [Input("operator", required: true)]
        public Input<string> Operator { get; set; } = null!;

        public SourceBingAdsConfigurationAccountNameArgs()
        {
        }
        public static new SourceBingAdsConfigurationAccountNameArgs Empty => new SourceBingAdsConfigurationAccountNameArgs();
    }
}
