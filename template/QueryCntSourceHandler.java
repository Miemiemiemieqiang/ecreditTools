package com.ecreditpal.vodka.variable.handler.source;

import com.ecreditpal.vodka.variable.feign.entity.response.advance.AdvanceMultiPlatformDetectionResp;
import com.ecreditpal.vodka.variable.feign.entity.response.advance.StatisticCustomerInfo;
import com.ecreditpal.vodka.variable.handler.Invoke;
import com.ecreditpal.vodka.variable.handler.Source;
import com.ecreditpal.vodka.variable.handler.SourceHandler;

import java.util.List;
import java.util.Optional;


/**
 * Created by {{.Author}} on {{.Datetime}}
 */
@Source(value = "{{.Lower}}", alias = "{{.Lower}}")
public class {{.Upper}}SourceHandler extends SourceHandler {

    @Invoke
    public Integer getSource(AdvanceMultiPlatformDetectionResp resp) {
        if (resp == null) return null;
        List<StatisticCustomerInfo> statisticCustomerInfo = resp.getStatistics().getStatisticCustomerInfo();
        Optional<StatisticCustomerInfo> any = statisticCustomerInfo.stream().filter(info -> "1-90d".equals(info.getTimePeriod())).findAny();
        return any.map(StatisticCustomerInfo::getQueryCount).orElse(null);
    }
}
