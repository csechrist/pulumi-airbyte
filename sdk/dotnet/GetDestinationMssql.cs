// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte
{
    public static class GetDestinationMssql
    {
        /// <summary>
        /// DestinationMssql DataSource
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
        ///     var myDestinationMssql = Airbyte.GetDestinationMssql.Invoke(new()
        ///     {
        ///         DestinationId = "...my_destination_id...",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetDestinationMssqlResult> InvokeAsync(GetDestinationMssqlArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDestinationMssqlResult>("airbyte:index/getDestinationMssql:getDestinationMssql", args ?? new GetDestinationMssqlArgs(), options.WithDefaults());

        /// <summary>
        /// DestinationMssql DataSource
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
        ///     var myDestinationMssql = Airbyte.GetDestinationMssql.Invoke(new()
        ///     {
        ///         DestinationId = "...my_destination_id...",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetDestinationMssqlResult> Invoke(GetDestinationMssqlInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDestinationMssqlResult>("airbyte:index/getDestinationMssql:getDestinationMssql", args ?? new GetDestinationMssqlInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDestinationMssqlArgs : global::Pulumi.InvokeArgs
    {
        [Input("destinationId", required: true)]
        public string DestinationId { get; set; } = null!;

        public GetDestinationMssqlArgs()
        {
        }
        public static new GetDestinationMssqlArgs Empty => new GetDestinationMssqlArgs();
    }

    public sealed class GetDestinationMssqlInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("destinationId", required: true)]
        public Input<string> DestinationId { get; set; } = null!;

        public GetDestinationMssqlInvokeArgs()
        {
        }
        public static new GetDestinationMssqlInvokeArgs Empty => new GetDestinationMssqlInvokeArgs();
    }


    [OutputType]
    public sealed class GetDestinationMssqlResult
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
        private GetDestinationMssqlResult(
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