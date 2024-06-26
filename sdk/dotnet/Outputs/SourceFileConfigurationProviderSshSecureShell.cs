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
    public sealed class SourceFileConfigurationProviderSshSecureShell
    {
        public readonly string Host;
        public readonly string? Password;
        /// <summary>
        /// Default: "22"
        /// </summary>
        public readonly string? Port;
        public readonly string User;

        [OutputConstructor]
        private SourceFileConfigurationProviderSshSecureShell(
            string host,

            string? password,

            string? port,

            string user)
        {
            Host = host;
            Password = password;
            Port = port;
            User = user;
        }
    }
}
