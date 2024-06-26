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
    public sealed class DestinationClickhouseConfigurationTunnelMethod
    {
        public readonly Outputs.DestinationClickhouseConfigurationTunnelMethodNoTunnel? NoTunnel;
        public readonly Outputs.DestinationClickhouseConfigurationTunnelMethodPasswordAuthentication? PasswordAuthentication;
        public readonly Outputs.DestinationClickhouseConfigurationTunnelMethodSshKeyAuthentication? SshKeyAuthentication;

        [OutputConstructor]
        private DestinationClickhouseConfigurationTunnelMethod(
            Outputs.DestinationClickhouseConfigurationTunnelMethodNoTunnel? noTunnel,

            Outputs.DestinationClickhouseConfigurationTunnelMethodPasswordAuthentication? passwordAuthentication,

            Outputs.DestinationClickhouseConfigurationTunnelMethodSshKeyAuthentication? sshKeyAuthentication)
        {
            NoTunnel = noTunnel;
            PasswordAuthentication = passwordAuthentication;
            SshKeyAuthentication = sshKeyAuthentication;
        }
    }
}
