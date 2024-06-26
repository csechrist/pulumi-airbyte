// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte
{
    public static class GetDestinationLangchain
    {
        /// <summary>
        /// DestinationLangchain DataSource
        /// 
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Airbyte = Pulumi.Airbyte;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var myDestinationLangchain = Airbyte.GetDestinationLangchain.Invoke(new()
        ///     {
        ///         DestinationId = "...my_destination_id...",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetDestinationLangchainResult> InvokeAsync(GetDestinationLangchainArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDestinationLangchainResult>("airbyte:index/getDestinationLangchain:getDestinationLangchain", args ?? new GetDestinationLangchainArgs(), options.WithDefaults());

        /// <summary>
        /// DestinationLangchain DataSource
        /// 
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Airbyte = Pulumi.Airbyte;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var myDestinationLangchain = Airbyte.GetDestinationLangchain.Invoke(new()
        ///     {
        ///         DestinationId = "...my_destination_id...",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetDestinationLangchainResult> Invoke(GetDestinationLangchainInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDestinationLangchainResult>("airbyte:index/getDestinationLangchain:getDestinationLangchain", args ?? new GetDestinationLangchainInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDestinationLangchainArgs : global::Pulumi.InvokeArgs
    {
        [Input("destinationId", required: true)]
        public string DestinationId { get; set; } = null!;

        public GetDestinationLangchainArgs()
        {
        }
        public static new GetDestinationLangchainArgs Empty => new GetDestinationLangchainArgs();
    }

    public sealed class GetDestinationLangchainInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("destinationId", required: true)]
        public Input<string> DestinationId { get; set; } = null!;

        public GetDestinationLangchainInvokeArgs()
        {
        }
        public static new GetDestinationLangchainInvokeArgs Empty => new GetDestinationLangchainInvokeArgs();
    }


    [OutputType]
    public sealed class GetDestinationLangchainResult
    {
        /// <summary>
        /// The values required to configure the destination. Parsed as JSON.
        /// </summary>
        public readonly string Configuration;
        public readonly string DestinationId;
        public readonly string DestinationType;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;
        public readonly string WorkspaceId;

        [OutputConstructor]
        private GetDestinationLangchainResult(
            string configuration,

            string destinationId,

            string destinationType,

            string id,

            string name,

            string workspaceId)
        {
            Configuration = configuration;
            DestinationId = destinationId;
            DestinationType = destinationType;
            Id = id;
            Name = name;
            WorkspaceId = workspaceId;
        }
    }
}
