// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte
{
    public static class GetSourceSalesloft
    {
        /// <summary>
        /// SourceSalesloft DataSource
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
        ///     var mySourceSalesloft = Airbyte.GetSourceSalesloft.Invoke(new()
        ///     {
        ///         SourceId = "...my_source_id...",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetSourceSalesloftResult> InvokeAsync(GetSourceSalesloftArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSourceSalesloftResult>("airbyte:index/getSourceSalesloft:getSourceSalesloft", args ?? new GetSourceSalesloftArgs(), options.WithDefaults());

        /// <summary>
        /// SourceSalesloft DataSource
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
        ///     var mySourceSalesloft = Airbyte.GetSourceSalesloft.Invoke(new()
        ///     {
        ///         SourceId = "...my_source_id...",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetSourceSalesloftResult> Invoke(GetSourceSalesloftInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSourceSalesloftResult>("airbyte:index/getSourceSalesloft:getSourceSalesloft", args ?? new GetSourceSalesloftInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSourceSalesloftArgs : global::Pulumi.InvokeArgs
    {
        [Input("sourceId", required: true)]
        public string SourceId { get; set; } = null!;

        public GetSourceSalesloftArgs()
        {
        }
        public static new GetSourceSalesloftArgs Empty => new GetSourceSalesloftArgs();
    }

    public sealed class GetSourceSalesloftInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("sourceId", required: true)]
        public Input<string> SourceId { get; set; } = null!;

        public GetSourceSalesloftInvokeArgs()
        {
        }
        public static new GetSourceSalesloftInvokeArgs Empty => new GetSourceSalesloftInvokeArgs();
    }


    [OutputType]
    public sealed class GetSourceSalesloftResult
    {
        /// <summary>
        /// The values required to configure the source. Parsed as JSON.
        /// </summary>
        public readonly string Configuration;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;
        public readonly string SourceId;
        public readonly string SourceType;
        public readonly string WorkspaceId;

        [OutputConstructor]
        private GetSourceSalesloftResult(
            string configuration,

            string id,

            string name,

            string sourceId,

            string sourceType,

            string workspaceId)
        {
            Configuration = configuration;
            Id = id;
            Name = name;
            SourceId = sourceId;
            SourceType = sourceType;
            WorkspaceId = workspaceId;
        }
    }
}
