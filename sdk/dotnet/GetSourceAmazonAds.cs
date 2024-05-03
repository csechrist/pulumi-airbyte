// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte
{
    public static class GetSourceAmazonAds
    {
        /// <summary>
        /// SourceAmazonAds DataSource
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
        ///     var mySourceAmazonads = Airbyte.GetSourceAmazonAds.Invoke(new()
        ///     {
        ///         SourceId = "...my_source_id...",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetSourceAmazonAdsResult> InvokeAsync(GetSourceAmazonAdsArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSourceAmazonAdsResult>("airbyte:index/getSourceAmazonAds:getSourceAmazonAds", args ?? new GetSourceAmazonAdsArgs(), options.WithDefaults());

        /// <summary>
        /// SourceAmazonAds DataSource
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
        ///     var mySourceAmazonads = Airbyte.GetSourceAmazonAds.Invoke(new()
        ///     {
        ///         SourceId = "...my_source_id...",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetSourceAmazonAdsResult> Invoke(GetSourceAmazonAdsInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSourceAmazonAdsResult>("airbyte:index/getSourceAmazonAds:getSourceAmazonAds", args ?? new GetSourceAmazonAdsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSourceAmazonAdsArgs : global::Pulumi.InvokeArgs
    {
        [Input("sourceId", required: true)]
        public string SourceId { get; set; } = null!;

        public GetSourceAmazonAdsArgs()
        {
        }
        public static new GetSourceAmazonAdsArgs Empty => new GetSourceAmazonAdsArgs();
    }

    public sealed class GetSourceAmazonAdsInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("sourceId", required: true)]
        public Input<string> SourceId { get; set; } = null!;

        public GetSourceAmazonAdsInvokeArgs()
        {
        }
        public static new GetSourceAmazonAdsInvokeArgs Empty => new GetSourceAmazonAdsInvokeArgs();
    }


    [OutputType]
    public sealed class GetSourceAmazonAdsResult
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
        private GetSourceAmazonAdsResult(
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