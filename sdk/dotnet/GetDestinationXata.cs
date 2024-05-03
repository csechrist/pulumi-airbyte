// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte
{
    public static class GetDestinationXata
    {
        /// <summary>
        /// DestinationXata DataSource
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
        ///     var myDestinationXata = Airbyte.GetDestinationXata.Invoke(new()
        ///     {
        ///         DestinationId = "...my_destination_id...",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetDestinationXataResult> InvokeAsync(GetDestinationXataArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDestinationXataResult>("airbyte:index/getDestinationXata:getDestinationXata", args ?? new GetDestinationXataArgs(), options.WithDefaults());

        /// <summary>
        /// DestinationXata DataSource
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
        ///     var myDestinationXata = Airbyte.GetDestinationXata.Invoke(new()
        ///     {
        ///         DestinationId = "...my_destination_id...",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetDestinationXataResult> Invoke(GetDestinationXataInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDestinationXataResult>("airbyte:index/getDestinationXata:getDestinationXata", args ?? new GetDestinationXataInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDestinationXataArgs : global::Pulumi.InvokeArgs
    {
        [Input("destinationId", required: true)]
        public string DestinationId { get; set; } = null!;

        public GetDestinationXataArgs()
        {
        }
        public static new GetDestinationXataArgs Empty => new GetDestinationXataArgs();
    }

    public sealed class GetDestinationXataInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("destinationId", required: true)]
        public Input<string> DestinationId { get; set; } = null!;

        public GetDestinationXataInvokeArgs()
        {
        }
        public static new GetDestinationXataInvokeArgs Empty => new GetDestinationXataInvokeArgs();
    }


    [OutputType]
    public sealed class GetDestinationXataResult
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
        private GetDestinationXataResult(
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