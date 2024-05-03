// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte
{
    public static class GetDestinationPubsub
    {
        /// <summary>
        /// DestinationPubsub DataSource
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
        ///     var myDestinationPubsub = Airbyte.GetDestinationPubsub.Invoke(new()
        ///     {
        ///         DestinationId = "...my_destination_id...",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetDestinationPubsubResult> InvokeAsync(GetDestinationPubsubArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDestinationPubsubResult>("airbyte:index/getDestinationPubsub:getDestinationPubsub", args ?? new GetDestinationPubsubArgs(), options.WithDefaults());

        /// <summary>
        /// DestinationPubsub DataSource
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
        ///     var myDestinationPubsub = Airbyte.GetDestinationPubsub.Invoke(new()
        ///     {
        ///         DestinationId = "...my_destination_id...",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetDestinationPubsubResult> Invoke(GetDestinationPubsubInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDestinationPubsubResult>("airbyte:index/getDestinationPubsub:getDestinationPubsub", args ?? new GetDestinationPubsubInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDestinationPubsubArgs : global::Pulumi.InvokeArgs
    {
        [Input("destinationId", required: true)]
        public string DestinationId { get; set; } = null!;

        public GetDestinationPubsubArgs()
        {
        }
        public static new GetDestinationPubsubArgs Empty => new GetDestinationPubsubArgs();
    }

    public sealed class GetDestinationPubsubInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("destinationId", required: true)]
        public Input<string> DestinationId { get; set; } = null!;

        public GetDestinationPubsubInvokeArgs()
        {
        }
        public static new GetDestinationPubsubInvokeArgs Empty => new GetDestinationPubsubInvokeArgs();
    }


    [OutputType]
    public sealed class GetDestinationPubsubResult
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
        private GetDestinationPubsubResult(
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