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
    public sealed class DestinationAwsDatalakeConfigurationCredentials
    {
        public readonly Outputs.DestinationAwsDatalakeConfigurationCredentialsIamRole? IamRole;
        public readonly Outputs.DestinationAwsDatalakeConfigurationCredentialsIamUser? IamUser;

        [OutputConstructor]
        private DestinationAwsDatalakeConfigurationCredentials(
            Outputs.DestinationAwsDatalakeConfigurationCredentialsIamRole? iamRole,

            Outputs.DestinationAwsDatalakeConfigurationCredentialsIamUser? iamUser)
        {
            IamRole = iamRole;
            IamUser = iamUser;
        }
    }
}
