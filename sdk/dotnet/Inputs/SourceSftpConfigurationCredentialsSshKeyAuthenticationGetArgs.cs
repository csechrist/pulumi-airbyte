// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class SourceSftpConfigurationCredentialsSshKeyAuthenticationGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("authSshKey", required: true)]
        private Input<string>? _authSshKey;

        /// <summary>
        /// OS-level user account ssh key credentials in RSA PEM format ( created with ssh-keygen -t rsa -m PEM -f myuser_rsa )
        /// </summary>
        public Input<string>? AuthSshKey
        {
            get => _authSshKey;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _authSshKey = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        public SourceSftpConfigurationCredentialsSshKeyAuthenticationGetArgs()
        {
        }
        public static new SourceSftpConfigurationCredentialsSshKeyAuthenticationGetArgs Empty => new SourceSftpConfigurationCredentialsSshKeyAuthenticationGetArgs();
    }
}
