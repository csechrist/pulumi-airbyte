// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte
{
    public static class GetSourcePexelsApi
    {
        /// <summary>
        /// SourcePexelsAPI DataSource
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
        ///     var mySourcePexelsapi = Airbyte.GetSourcePexelsApi.Invoke(new()
        ///     {
        ///         SourceId = "...my_source_id...",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetSourcePexelsApiResult> InvokeAsync(GetSourcePexelsApiArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSourcePexelsApiResult>("airbyte:index/getSourcePexelsApi:getSourcePexelsApi", args ?? new GetSourcePexelsApiArgs(), options.WithDefaults());

        /// <summary>
        /// SourcePexelsAPI DataSource
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
        ///     var mySourcePexelsapi = Airbyte.GetSourcePexelsApi.Invoke(new()
        ///     {
        ///         SourceId = "...my_source_id...",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetSourcePexelsApiResult> Invoke(GetSourcePexelsApiInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSourcePexelsApiResult>("airbyte:index/getSourcePexelsApi:getSourcePexelsApi", args ?? new GetSourcePexelsApiInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSourcePexelsApiArgs : global::Pulumi.InvokeArgs
    {
        [Input("sourceId", required: true)]
        public string SourceId { get; set; } = null!;

        public GetSourcePexelsApiArgs()
        {
        }
        public static new GetSourcePexelsApiArgs Empty => new GetSourcePexelsApiArgs();
    }

    public sealed class GetSourcePexelsApiInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("sourceId", required: true)]
        public Input<string> SourceId { get; set; } = null!;

        public GetSourcePexelsApiInvokeArgs()
        {
        }
        public static new GetSourcePexelsApiInvokeArgs Empty => new GetSourcePexelsApiInvokeArgs();
    }


    [OutputType]
    public sealed class GetSourcePexelsApiResult
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
        private GetSourcePexelsApiResult(
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