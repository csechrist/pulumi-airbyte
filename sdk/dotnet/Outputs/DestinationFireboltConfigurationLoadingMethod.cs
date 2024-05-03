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
    public sealed class DestinationFireboltConfigurationLoadingMethod
    {
        public readonly Outputs.DestinationFireboltConfigurationLoadingMethodExternalTableViaS3? ExternalTableViaS3;
        public readonly Outputs.DestinationFireboltConfigurationLoadingMethodSqlInserts? SqlInserts;

        [OutputConstructor]
        private DestinationFireboltConfigurationLoadingMethod(
            Outputs.DestinationFireboltConfigurationLoadingMethodExternalTableViaS3? externalTableViaS3,

            Outputs.DestinationFireboltConfigurationLoadingMethodSqlInserts? sqlInserts)
        {
            ExternalTableViaS3 = externalTableViaS3;
            SqlInserts = sqlInserts;
        }
    }
}